# Fixes applied to Attio's OpenAPI spec before code generation.
walk(
  if type == "object" then
    # Attio's oneOf variants overlap, so generated oneOf decoders reject values
    # that match more than one variant. anyOf takes the first match instead.
    (if has("oneOf") then .anyOf = .oneOf | del(.oneOf) else . end)
    # `additionalProperties: {}` or `true` means the same as leaving it out, but
    # inside anyOf it makes openapi-generator emit a Go field that does not compile.
    | (if .additionalProperties == {} or .additionalProperties == true then del(.additionalProperties) else . end)
  else . end
)
