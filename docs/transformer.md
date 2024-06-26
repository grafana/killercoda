# About the transformer tool

The transformer tool generates Killercoda tutorials from Grafana documentation source files.

To use the transformer tool, you need to add Killercoda metadata to the source file front matter and annotate your source file with Killercoda directives.

## Metadata

You specify Killercoda tutorial metadata in the source file front matter as the value for the `killercoda` field.
The tool uses the metadata to perform preprocessing on the source file and generate the Killercoda configuration files for the tutorial.

| Field                                    | Type   | Description                                                                                                                                                          |
| ---------------------------------------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `killercoda.backend.imageid`             | String | The name of the Killercoda environment's backend image. Supported values include `ubuntu`.                                                                           |
| `killercoda.description`                 | String | The description displayed on the Killercoda website                                                                                                                  |
| `killercoda.details.finish.text`         | String | The filename of the finish page Markdown source in the grafana/killercoda repository. A [finish directive](#finish) in the documentation source overrides this.      |
| `killercoda.details.intro.text`          | String | The filename of the introduction page Markdown source in the grafana/killercoda repository. An [intro directive](#intro) in the documentation source overrides this. |
| `killercoda.preprocessing.substitutions` | Array  | Substitute matches of a regular expression with a replacement. For more information, refer to [Substitutions](#substitutions).                                       |
| `killercoda.title`                       | String | The title for the tutorial on the Killercoda website.                                                                                                                |


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
- [Define an introduction page](#intro)
- [Define a finish page](#finish)
- [Define step pages](#step)
- [Ignore parts of the documentation](#ignore)
- [Include extra parts not in the website page](#include)

### Copy

Copy directives tell the transform tool to make the contained fenced code block copyable.

The start marker is:

```markdown
<!-- INTERACTIVE copy START -->
```

The end marker is:

```markdown
<!-- INTERACTIVE copy END -->
```

#### Examples

````markdown
<!-- INTERACTIVE copy START -->

```bash
echo 'Hello, world!'
```

<!-- INTERACTIVE copy END -->
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
<!-- INTERACTIVE exec START -->
```

The end marker is:

```markdown
<!-- INTERACTIVE exec END -->
```

#### Examples

````markdown
<!-- INTERACTIVE exec START -->

```bash
echo 'Hello, world!'
```

<!-- INTERACTIVE exec END -->
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
If this is present, it overrides the `killercoda.details.finish.text` front matter.

The start marker is:

```markdown
<!-- INTERACTIVE finish.md START -->
```

The end marker is:

```markdown
<!-- INTERACTIVE finish.md END -->
```

### Ignore

The ignore directive tells the transform tool to skip the contents within the markers when generating the Killercoda page.

The start marker is:

```markdown
<!-- INTERACTIVE ignore START -->
```

The end marker is:

```markdown
<!-- INTERACTIVE ignore END -->
```

#### Examples

```markdown
Information common to both pages.

<!-- INTERACTIVE ignore START -->

Information unique to the Grafana website page.

<!-- INTERACTIVE ignore END -->
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
<!-- INTERACTIVE include START -->
```

The end marker is:

```markdown
<!-- INTERACTIVE include END -->
```

#### Examples

```markdown
Information common to both pages.

<!-- INTERACTIVE include START -->
<!-- Information unique to the Killercoda page. -->
<!-- INTERACTIVE include END -->
```

Produces:

```markdown
Information common to both pages.

Information unique to the Killercoda page.
```

### Intro

The intro directive specifies the start and end of the section of the file to use as the Killercoda introduction page.
If this is present, it overrides the `killercoda.details.intro.text` front matter.

The start marker is:

```markdown
<!-- INTERACTIVE intro.md START -->
```

The end marker is:

```markdown
<!-- INTERACTIVE intro.md END -->
```

### Step

The step directive specifies the start and end of the section of the file to use as a Killercoda step page.

The start marker is the following, where _`<N>`_ is the number of the step:

```markdown
<!-- INTERACTIVE step<N>.md START -->
```

The end marker is the following, where _`<N>`_ is the number of the step:

```markdown
<!-- INTERACTIVE step<N>.md END -->
```

## Generate a tutorial

You can generate a new tutorial from existing documentation using the transformer tool.
After generation, a [GitHub Actions workflow](./.github/workflows/regenerate-tutorials) updates the tutorial when the documentation source changes.

### Before you begin

1. Clone the repository with the source documentation.
1. Clone the Killercoda repository.
1. [Download and install Go](https://go.dev/doc/install) to run the transformer tool.

To generate a tutorial:

1. In the source repository, check out a new branch.

   Give the branch a name that reflects the planned changes.
   For example, `2024-06-killercoda-tutorial-for-loki-quickstart`.

1. In the Killercoda repository, check out a new branch.

   For convenience, name the branch the same as you did in the source repository.

1. In the source repository, run `make docs` from the `docs/` directory so that you can verify your changes don't break the rendered documentation.

1. In the source file, add the Killercoda metadata to the front matter.

   Front matter is YAML metadata written before the page's content.
   For more information, refer to the [Hugo front matter documentation](https://gohugo.io/content-management/front-matter/).

   For an example Killercoda front matter, refer to [Metadata](#Metadata)

1. Configure an introduction page.

   Use one of the two options:

   1. In the source repository, add [intro](#intro) directives.
   1. 1. In the Killercoda repository, add a `intro.md` file in the output tutorial directory.
      1. In the source file, add the `killercoda.details.intro.text` field with the value `intro.md`.

1. Add directives for each step in the tutorial.

   Each step starts at the [step](#step) directive start marker and ends at the [step](#step) directive end marker.
   Include at least one step.
   The first step use the start marker `<!-- INTERACTIVE step1.md START -->` and the end marker `<!-- INTERACTIVE step1.md END -->`

1. Configure a finish page.

   Use one of the two options:

   1. In the source repository, add [finish](#finish) directives.
   1. 1. In the Killercoda repository, add a `finish.md` file in the output tutorial directory.
      1. In the source file, add the `killercoda.details.finish.text` field with the value `finish.md`.

1. Generate the tutorial.

   1. In the Killercoda repository, change to the `tools/transformer` directory.
   1. Run `go run ./ <PATH TO SOURCE FILE> <PATH TO OUTPUT DIRECTORY>`

      - _`<PATH TO SOURCE FILE>`_ is the path to the documentation file.

        For example, `~/ext/grafana/loki/docs/sources/get-started/quick-start.md`.

      - _`<PATH TO OUTPUT DIRECTORY>`_ is the path to the output directory in the Killercoda repository.

        For example, `~/ext/grafana/killercoda/loki/loki-quickstart`.

      The tool creates the `index.json` and step Markdown files.

1. Add the tutorial to the Killercoda structure file.

   The structure file is a `structure.json` file in each directory that contains tutorials.

   The Loki structure file is [`loki/structure.json`](./loki/structure.json).

1. In each repository, commit your changes, push your branch, and open a pull request.

1. A Killercoda maintainer reviews the PR to ensure that the generate tutorial works as expected.

## Scripts and extra course files

If your tutorial requires scripts or extra files, make sure to manually add them to the tutorial directory in the Killercoda repository. For example, if your tutorial requires a script to run:

1. Add the bash script to the tutorial directory in the Killercoda repository. Refer to the [what-is-loki](../loki/what-is-loki/) tutorial for an example.
2. Add the script to the desired step within the `index.json` file. Note that `foreground` scripts run in the foreground (seen in terminal), and `background` scripts run in the background (run in background thread). For example:
   ```json
   {
      "title": "What is Loki?",
      "description": "A sandbox enviroment to introduce Loki to new users.",
      "details": {
         "intro": {
         "text": "intro.md",
         "foreground": "script1.sh"
         },
         "steps": [
         {
            "text": "step1.md",
            "foreground": "script2.sh"
         }
         ],
         "finish": {
         "text": "finished.md"
         }
      },
      "backend": {
         "imageid": "ubuntu"
      }
   }

   ```

For extra assets, such as images or configuration files:
1. Create a directory called `assets` in the tutorial directory.
2. Add the assets to the `assets` directory.
3. Add the mount path to the `index.json` file:
   ```json
   {
      "title": "Grafana Basics",
      "description": "In this demo learn how to install and configure Grafana",
      "details": {
         "intro": {
         "text": "intro.md"
         },
         "steps": [],
         "finish": {
         "text": "finished.md"
         },
         "assets": {
         "host01": [
            {"file": "*", "target": "/education"}
         ]
         }
      },
      "backend": {
         "imageid": "ubuntu"
      }
   }
   ```
   Refer to the [grafana-basics](../grafana/grafana-basics/index.json) tutorial for an example.