# goddlgen
A tool which generates Database DDL from input definitions.

## Descriptor
The application requires a descriptor to guide the generation process.  The descriptor defines meta-data, input files and output files to be generated.

### Descriptor Definition
The following table describes the elements of the Descriptor file in more detail.

| Key Name                   | Description                                                                                                                         |
|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| input.json.folderName      | The name of the folder which contains JSON model data to read for generation.                                                       |
| output.sql.folderName      | The name of the folder to which the DDL will be generated.                                                                          |

### Sample Tool Invocations
The following example is a typical use case generating everything defined by the descriptor:

        ddlgen generate -d descriptor.yml

### Sample Descriptor

```yaml
input:
  json:
    folderName: "./sample_model/small"

output:
  sql:
    folderName: "./sample_output/small"
```
