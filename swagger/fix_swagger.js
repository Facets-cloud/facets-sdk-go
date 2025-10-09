#!/usr/bin/env node

const fs = require('fs');

// Get command line arguments
const args = process.argv.slice(2);
if (args.length < 2) {
    console.error('Usage: node fix_swagger.js <input-file> <output-file>');
    process.exit(1);
}

const inputFile = args[0];
const outputFile = args[1];

console.log(`Reading OpenAPI spec from: ${inputFile}`);

// Read the OpenAPI spec
let spec;
try {
    const content = fs.readFileSync(inputFile, 'utf8');
    spec = JSON.parse(content);
} catch (error) {
    console.error(`Error reading input file: ${error.message}`);
    process.exit(1);
}

console.log(`OpenAPI version: ${spec.openapi}`);

// Convert OpenAPI 3.0 to Swagger 2.0 format
if (spec.openapi && spec.openapi.startsWith('3.')) {
    console.log('Converting OpenAPI 3.0 to Swagger 2.0...');

    const swagger2 = {
        swagger: '2.0',
        info: spec.info || {},
        host: '',
        basePath: '/',
        schemes: ['https'],
        paths: spec.paths || {},
        definitions: {},
        securityDefinitions: {}
    };

    // Extract host from servers
    if (spec.servers && spec.servers.length > 0) {
        const serverUrl = new URL(spec.servers[0].url);
        swagger2.host = serverUrl.host;
        swagger2.basePath = serverUrl.pathname === '/' ? '/' : serverUrl.pathname;
    }

    // Convert components/schemas to definitions
    if (spec.components && spec.components.schemas) {
        swagger2.definitions = spec.components.schemas;
    }

    // Convert security schemes
    if (spec.components && spec.components.securitySchemes) {
        for (const [name, scheme] of Object.entries(spec.components.securitySchemes)) {
            if (scheme.type === 'http' && scheme.scheme === 'basic') {
                swagger2.securityDefinitions[name] = {
                    type: 'basic'
                };
            } else if (scheme.type === 'http' && scheme.scheme === 'bearer') {
                swagger2.securityDefinitions[name] = {
                    type: 'apiKey',
                    name: 'Authorization',
                    in: 'header'
                };
            } else if (scheme.type === 'apiKey') {
                swagger2.securityDefinitions[name] = {
                    type: 'apiKey',
                    name: scheme.name,
                    in: scheme.in
                };
            }
        }
    }

    // Convert paths - handle OpenAPI 3.0 specific features
    for (const [pathName, pathItem] of Object.entries(swagger2.paths)) {
        for (const [method, operation] of Object.entries(pathItem)) {
            if (typeof operation !== 'object' || !operation) continue;
            if (['parameters', 'servers', '$ref'].includes(method)) continue;

            // Fix and validate parameters
            if (operation.parameters) {
                operation.parameters = operation.parameters.filter(param => {
                    // Remove invalid parameters
                    if (!param.name || !param.in) return false;

                    // Fix schema references in parameters
                    if (param.schema) {
                        if (param.schema.$ref) {
                            param.schema.$ref = param.schema.$ref.replace('#/components/schemas/', '#/definitions/');
                        }
                        // Remove invalid fields from parameter schemas
                        delete param.schema.nullable;
                        delete param.schema.example;
                    }

                    // Ensure required fields
                    if (param.in === 'body' && !param.schema) {
                        param.schema = { type: 'object' };
                    }

                    return true;
                });
            }

            // Convert requestBody to body parameter
            if (operation.requestBody) {
                const requestBody = operation.requestBody;
                const content = requestBody.content;

                if (content) {
                    const contentType = Object.keys(content)[0];
                    const schema = content[contentType].schema;

                    if (!operation.parameters) {
                        operation.parameters = [];
                    }

                    operation.parameters.push({
                        in: 'body',
                        name: 'body',
                        required: requestBody.required || false,
                        description: requestBody.description || '',
                        schema: schema || { type: 'object' }
                    });

                    // Add consumes
                    if (!operation.consumes) {
                        operation.consumes = [contentType];
                    }
                }

                delete operation.requestBody;
            }

            // Convert responses
            if (operation.responses) {
                for (const [statusCode, response] of Object.entries(operation.responses)) {
                    if (response.content) {
                        const contentType = Object.keys(response.content)[0];
                        const mediaType = response.content[contentType];

                        response.schema = mediaType.schema || { type: 'object' };
                        delete response.content;

                        // Add produces
                        if (!operation.produces) {
                            operation.produces = [contentType];
                        }
                    }

                    // Ensure all responses have proper structure
                    if (!response.description) {
                        response.description = `Response ${statusCode}`;
                    }
                }
            }

            // Ensure operation has required fields
            if (!operation.responses) {
                operation.responses = {
                    '200': {
                        description: 'Success'
                    }
                };
            }
        }
    }

    // Add security if present
    if (spec.security) {
        swagger2.security = spec.security;
    }

    // Add tags if present
    if (spec.tags) {
        swagger2.tags = spec.tags;
    }

    spec = swagger2;
}

// Clean up $ref paths in definitions
if (spec.definitions) {
    const cleanRefs = (obj) => {
        if (typeof obj !== 'object' || obj === null) return;

        for (const key in obj) {
            if (key === '$ref' && typeof obj[key] === 'string') {
                // Convert #/components/schemas/X to #/definitions/X
                obj[key] = obj[key].replace('#/components/schemas/', '#/definitions/');
            } else if (typeof obj[key] === 'object') {
                cleanRefs(obj[key]);
            }
        }
    };

    cleanRefs(spec);
}

// Remove invalid or problematic fields
const removeInvalidFields = (obj) => {
    if (typeof obj !== 'object' || obj === null) return;

    for (const key in obj) {
        // Remove nullable as it's not valid in Swagger 2.0
        if (key === 'nullable') {
            delete obj[key];
        }
        // Remove OpenAPI 3.0 specific fields
        else if (key === 'oneOf' || key === 'anyOf' || key === 'allOf') {
            // Convert oneOf/anyOf/allOf to object type for Swagger 2.0 compatibility
            if (Array.isArray(obj[key]) && obj[key].length > 0) {
                // Use the first schema as a fallback
                const firstSchema = obj[key][0];
                if (firstSchema.$ref) {
                    obj.$ref = firstSchema.$ref;
                } else if (firstSchema.type) {
                    obj.type = firstSchema.type;
                    if (firstSchema.properties) obj.properties = firstSchema.properties;
                }
            }
            delete obj[key];
        }
        else if (typeof obj[key] === 'object') {
            removeInvalidFields(obj[key]);
        }
    }
};

removeInvalidFields(spec);

// Fix schema definitions - ensure all have valid types
if (spec.definitions) {
    for (const [name, schema] of Object.entries(spec.definitions)) {
        // If schema has no type and no $ref, default to object
        if (!schema.type && !schema.$ref && !schema.enum) {
            if (schema.properties) {
                schema.type = 'object';
            } else if (schema.items) {
                schema.type = 'array';
            }
        }

        // Fix array items that don't have type or $ref
        if (schema.type === 'array' && schema.items) {
            if (!schema.items.type && !schema.items.$ref && !schema.items.enum) {
                schema.items.type = 'object';
            }
        }

        // Remove empty or invalid properties
        if (schema.properties) {
            for (const [propName, propSchema] of Object.entries(schema.properties)) {
                if (!propSchema || typeof propSchema !== 'object') {
                    delete schema.properties[propName];
                } else if (!propSchema.type && !propSchema.$ref && !propSchema.enum && !propSchema.items) {
                    propSchema.type = 'string'; // Default to string
                }
            }
        }
    }
}

// Write the cleaned spec
try {
    fs.writeFileSync(outputFile, JSON.stringify(spec, null, 2), 'utf8');
    console.log(`Cleaned spec written to: ${outputFile}`);
    console.log('✓ OpenAPI spec cleaned successfully');
} catch (error) {
    console.error(`Error writing output file: ${error.message}`);
    process.exit(1);
}
