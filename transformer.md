# Transformer

The transformer tool generates Killercoda tutorials from Grafana documentation source files.

To use the transformer tool, you need to add Killercoda metadata to the source file front matter and annotate your source file with Killercoda directives.

## Metadata

You specify Killercoda tutorial metadata in the source file front matter as the value for the `killercoda` field.
The tool uses the metadata to perform preprocessing on the source file and generate the Killercoda configuration files for the tutorial.

| Field                                    | Type   | Description                                                                                                                    |
| ---------------------------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `killercoda.backend.imageid`             | String | The name of the Killercoda environment's backend image. Supported values include `ubuntu`.                                     |
| `killercoda.description`                 | String | The description displayed on the Killercoda website                                                                            |
| `killercoda.details.finish.text`         | String | The filename of the finish page Markdown source in the grafana/killercoda repository.                                          |
| `killercoda.details.intro.text`          | String | The filename of the introduction page Markdown source in the grafana/killercoda repository.                                    |
| `killercoda.preprocessing.substitutions` | Array  | Substitute matches of a regular expression with a replacement. For more information, refer to [Substitutions](#substitutions). |
| `killercoda.title`                       | String | The title for the tutorial on the Killercoda website.                                                                          |

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
      text: finished.md
  backend:
    imageid: ubuntu
```

### Substitutions

Substitutions substitutes all matches of a regular expression in a source file with a replacement.
Each substitution has two fields:

- `regexp`: An RE2 regular expression matched and `replacement`.
- `replacement`: The string that replaces the match. You can reference numbered capture groups using the `$` syntax.
  To reference the first capture group, use `$1`.

## Directives

Directives in the source file modify how the transformer tool generates the tutorial.
You write directives in the source file with HTML comments.

Use directives to:

- [Configure copyable code blocks](#copy)
- [Configure executable code blocks](#exec)
- [Define a finish page](#finish)
- [Define an introduction page](#intro)
- [Define step pages](#steps)
- [Ignore parts of the documentation](#ignore)
- [Include extra parts not in the website page](#include)

### Copy

Copy directives tell the transform tool to make the contained fenced code block copyable.

The start marker is:

```markdown
<!-- Killercoda copy START -->
```

The end marker is:

```markdown
<!-- Killercoda copy END -->
```

#### Examples

````markdown
<!-- Killercoda copy START -->

```bash
echo 'Hello, world!'
```

<!-- Killercoda copy END -->
````

Produces:

<!-- prettier-ignore-start -->

````markdown
```bash
echo 'Hello, world!'
```{{exec}}
````

<!-- prettier-ignore-end -->

### Exec

Exec directives tell the transform tool to make the contained fenced code block executable.

The start marker is:

```markdown
<!-- Killercoda exec START -->
```

The end marker is:

```markdown
<!-- Killercoda exec END -->
```

#### Examples

````markdown
<!-- Killercoda exec START -->

```bash
echo 'Hello, world!'
```

<!-- Killercoda exec END -->
````

Produces:

<!-- prettier-ignore-start -->

````markdown
```bash
echo 'Hello, world!'
```{{exec}}
````

<!-- prettier-ignore-end -->

### Finish

The finish directive specifies the start and end of the section of the file to use as the Killercoda finish page.

The start marker is:

```markdown
<!-- Killercoda finish.md START -->
```

The end marker is:

```markdown
<!-- Killercoda finish.md END -->
```

### Ignore

The ignore directive tells the transform tool to skip the contents within the markers when generating the Killercoda page.

The start marker is:

```markdown
<!-- Killercoda ignore START -->
```

The end marker is:

```markdown
<!-- Killercoda ignore END -->
```

#### Examples

```markdown
Information common to both pages.

<!-- Killercoda ignore START -->

Information unique to the Grafana website page.

<!-- Killercoda ignore END -->
```

Produces:

```markdown
Information common to both pages.
```

### Include

The include directive tells the transform tool to include the contents of the HTML comments within the markers as content when generating the Killercoda page.
The HTML comments aren't rendered on the Grafana website.

The start marker is:

```markdown
<!-- Killercoda include START -->
```

The end marker is:

```markdown
<!-- Killercoda include END -->
```

#### Examples

```markdown
Information common to both pages.

<!-- Killercoda include START -->
<!-- Information unique to the Killercoda page. -->
<!-- Killercoda include END -->
```

Produces:

```markdown
Information common to both pages.

Information unique to the Killercoda page.
```

### Intro

The intro directive specifies the start and end of the section of the file to use as the Killercoda introduction page.

The start marker is:

```markdown
<!-- Killercoda intro.md START -->
```

The end marker is:

```markdown
<!-- Killercoda intro.md END -->
```

### Steps

```markdown
<!-- Killercoda step<N>.md START -->
```

```markdown
<!-- Killercoda step<N>.md END -->
```
