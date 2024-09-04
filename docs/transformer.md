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

- [Configure executable code blocks](#exec)
- [Define pages](#page)
- [Ignore parts of the documentation](#ignore)

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

> #### NOTE
>
> By default, the tool makes `bash` fenced code blocks executable so you don't need `<!-- INTERACTIVE exec START/STOP -->` directives for bash code blocks.
> You can override this behavior with the `<!-- INTERACTIVE copy START/STOP -->` directives which take precedence over the default behavior.
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

> [!NOTE]
> By default, the tool makes all fenced code blocks other than `bash` copyable so you don't need `<!-- INTERACTIVE copy START/STOP -->` directives for those code blocks.
> You can override this behavior with the `<!-- INTERACTIVE exec START/STOP -->` directives which take precedence over the default behavior.

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
```{{copy}}
````

<!-- prettier-ignore-end -->


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

To do the inverse task, and ignore content in the website build, use the [`docs/ignore` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#docsignore).

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

### Page

The page directive tells the transform tool to use the content between the markers as the source for a Killercoda page.
The page's filename is the first argument to the directive.

Every tutorial must have at least the pages:

- `intro.md`: An introduction to the tutorial.
- `step1.md`: The first step in the tutorial.
- `finish.md`: A closing page that summarizes steps taken and includes next steps.

You can also add additional steps using the `step<N>.md`, where _`<N>`_ is in the range 2-20.
Steps must be sequential, you can't have `step1.md` and `step3.md` without a `step2.md`.

The start marker is:

```markdown
<!-- INTERACTIVE page <FILENAME> START -->
```

The end marker is:

```markdown
<!-- INTERACTIVE page <FILENAME> END -->
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

   For an example Killercoda front matter, refer to [Metadata](#metadata).

1. Configure an introduction page.

   Use one of the two options:

   1. In the source repository, add the [page](#page) directives with the `intro.md` argument.

      The start marker is `<!-- INTERACTIVE page intro.md START -->` and the end marker is `<!-- INTERACTIVE page intro.md END -->`.

   1. 1. In the Killercoda repository, add a `intro.md` file in the output tutorial directory.
      1. In the source file, add the `killercoda.details.intro.text` field with the value `intro.md`.

1. Add [page](#page) directives for each step in the tutorial.

   You must include at least one step.
   The first step uses the start marker `<!-- INTERACTIVE page step1.md START -->` and the end marker `<!-- INTERACTIVE page step1.md END -->`.

1. Configure a finish page.

   Use one of the two options:

   1. In the source repository, add the [page](#page) directives with the `finish.md` argument.

      The start marker is `<!-- INTERACTIVE page finish.md START -->` and the end marker is `<!-- INTERACTIVE page finish.md END -->`.

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

## Add foreground and background scripts

Foreground and background scripts are shell scripts that run during the introduction and finish pages of a tutorial.
The scripts are useful for setting up the environment for the tutorial and cleaning up after the tutorial.

Scripts are another asset that needs to be maintained, so use them sparingly.
A good example of using a script is to update the Docker Compose package before running the tutorial.

Use foreground scripts when you want to the user to see the script output.
Use background scripts when you want to hide the output of the script.

### Create your script

Since these scripts are primarily used for preparing the interactive environment, they are stored within the [`sandbox-scripts`](/sandbox-scripts/) directory in the Killercoda repository. Make sure to create your script in this directory and name with a unique name that reflects the purpose of the script. Note if a script already exists that performs the same function, you can reuse it.

Here is an example of a script that updates the Docker Compose package:

```bash
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install -y docker-compose-plugin && clear && echo "Setup complete. You may now begin the tutorial."
```

> [!TIP]
> Add a message at the end of the script to inform the user that the setup is complete.

### Add the script to the tutorial

To add the script to the tutorial, you need to add the script to the `killercoda` metadata in the source file. 

> [!NOTE]
> The transformer tool only supports foreground and background scripts for the introduction and finish pages.

The following example sets the foreground script for the introduction page to be `docker-compose-update.sh`:

```yaml
title: Quick start for Tempo
menuTitle: Quick start for Tempo
description: Use Docker to quickly view traces using K-6 and Tempo
weight: 600
killercoda:
  title: Quick start for Tempo
  description: Use Docker to quickly view traces using K-6 and Tempo
  details:
      intro:
         foreground: docker-compose-update.sh
  backend:
    imageid: ubuntu
```

In the example above, we have added a foreground script to the introduction page. The script is named `docker-compose-update.sh` and is located in the `sandbox-scripts` directory. The script will run when the introduction page is loaded.

The following example sets the foreground script for the introduction page to be `docker-compose-cleanup.sh`:

```yaml
title: Quick start for Tempo
menuTitle: Quick start for Tempo
description: Use Docker to quickly view traces using K-6 and Tempo
weight: 600
killercoda:
  title: Quick start for Tempo
  description: Use Docker to quickly view traces using K-6 and Tempo
  details:
      finish:
         background: docker-compose-cleanup.sh
  backend:
    imageid: ubuntu
```
