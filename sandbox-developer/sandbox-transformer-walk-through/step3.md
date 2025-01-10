# Learn the basic meta syntax

The Sandbox Transformer uses a special syntax to define the metadata for the course. This metadata is used to define the structure of the course and the content of each page.

Here is a breakdown of the basic meta syntax. For more information on the meta syntax, see the [full documentation](https://github.com/grafana/killercoda/blob/staging/docs/transformer.md)

## Metadata

You specify Killercoda tutorial metadata in the source file front matter as the value for the `killercoda`{{copy}} field.
The tool uses the metadata to perform preprocessing on the source file and generate the Killercoda configuration files for the tutorial. A table of the metadata fields can be found [here](https://github.com/grafana/killercoda/blob/staging/docs/transformer.md#metadata)

The following YAML demonstrates a number of the fields:

```yaml
killercoda:
  preprocessing:
    substitutions:
      - regexp: evaluate-loki-([^-]+)-
        replacement: evaluate-loki_${1}_
  title: Loki Quickstart Demo
  description: This sandbox provides an online enviroment for testing the Loki quickstart demo.
  details:
    finish:
      text: finish.md
  backend:
    imageid: ubuntu
```{{copy}}

## Directives

Directives in the source file modify how the transformer tool generates the tutorial.
You write directives in the source file with HTML comments.

### Page

The page directive tells the transform tool to use the content between the markers as the source for a Killercoda page.
The page’s filename is the first argument to the directive.

Every tutorial must have at least the pages:

- `intro.md`{{copy}}: An introduction to the tutorial.

- `step1.md`{{copy}}: The first step in the tutorial.

- `finish.md`{{copy}}: A closing page that summarizes steps taken and includes next steps.

You can also add additional steps using the `step<N>.md`{{copy}}, where _`<N>`{{copy}}_ is in the range 2-20.
Steps must be sequential, you can’t have `step1.md`{{copy}} and `step3.md`{{copy}} without a `step2.md`{{copy}}.

The start marker is:

```markdown
<!-- INTERACTIVE page <FILENAME> START -->
```{{copy}}

The end marker is:

```markdown
<!-- INTERACTIVE page <FILENAME> END -->
```{{copy}}

### Exec

Exec directives tell the transform tool to make the contained fenced code block executable.

> **Note:**
> By default, the tool makes `bash`{{copy}} fenced code blocks executable so you don’t need `<!-- INTERACTIVE exec START/STOP -->`{{copy}} directives for bash code blocks.
> You can override this behavior with the `<!-- INTERACTIVE copy START/STOP -->`{{copy}} directives which take precedence over the default behavior.

The start marker is:

```markdown
<!-- INTERACTIVE exec START -->
```{{copy}}

The end marker is:

```markdown
<!-- INTERACTIVE exec END -->
```{{copy}}

### Copy

Copy directives tell the transform tool to make the contained fenced code block copyable.

The start marker is:

```markdown
<!-- INTERACTIVE copy START -->
```{{copy}}

The end marker is:

```markdown
<!-- INTERACTIVE copy END -->
```{{copy}}

### Ignore

The ignore directive tells the transform tool to skip the contents within the markers when generating the Killercoda page.

The start marker is:

```markdown
<!-- INTERACTIVE ignore START -->
```{{copy}}

The end marker is:

```markdown
<!-- INTERACTIVE ignore END -->
```{{copy}}

## Examples

The best place to see how the meta syntax works is to look at the examples in the `docs/examples`{{copy}} directory of the `killercoda`{{copy}} repository. You can find examples of how to structure your markdown files and how to use the meta syntax to define the structure of your course.
