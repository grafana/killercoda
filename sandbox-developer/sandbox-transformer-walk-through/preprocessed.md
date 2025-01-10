---
title: Learn how to use the Sandbox Transformer
menuTitle: Learn how to use the Sandbox Transformer
description: Learn how to use the Sandbox Transformer to turn hugo docs into a course
weight: 250
killercoda:
  title: Learn how to use the Sandbox Transformer
  description: Learn how to use the Sandbox Transformer to turn hugo docs into a course
  backend:
    imageid: ubuntu
---

<!-- INTERACTIVE page intro.md START -->

# Learn how to use the Sandbox Transformer

The Sandbox Transformer is an experimental tool created by Grafana Labs to turn **Hugo Markdown** files into KillerCoda courses. This tool is still in development, but we`re excited to share it with you and get your feedback. In this tutorial, you will learn how to use the Sandbox Transformer to turn Hugo docs into a course.

> This tutorial will also work with basic Markdown files, however, there are certain Hugo specific features such as the document metadata which is required for the transformer to work. This my interfere with the rendering of the original Markdown file.

## What you will learn

- How to build the Sandbox Transformer
- Learn the basic meta syntax
- How to use the Sandbox Transformer to turn Hugo docs into a course

<!-- INTERACTIVE page intro.md END -->

<!-- INTERACTIVE page step1.md START -->

# Prerequisites

In this section we will cover the prerequisites you need to have in place in order to build and run the Sandbox Transformer.

## Clone the repository

First, you need to clone the repository to your local machine. You can do this by running the following command:

```bash
git clone https://github.com/grafana/killercoda.git
```

Its best practise to create a new branch for each new course you create. You can do this by running the following command:

```bash
git checkout -b my-new-course
```

## Install Go 

The Sandbox Transformer is written in Go, so you will need to have Go installed on your machine. You can download Go from the official website [here](https://golang.org/dl/). In this case we will install the Ubuntu package:

1. Download the Go package:
   
   ```bash
   wget https://go.dev/dl/go1.23.4.linux-386.tar.gz
   ```
1. Remove old versions of Go and Install the package:
   
   ```bash
   rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
   ```

1. Add the Go binary to your PATH:
   
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```
1. Verify the installation:
   
   ```bash
   go version
   ```

<!-- INTERACTIVE page step1.md END -->

<!-- INTERACTIVE page step2.md START -->

# Build the Sandbox Transformer

Now that you have the repository cloned and Go installed, you can build the Sandbox Transformer.

## Build the transformer

To build the transformer: 

1. navigate to the `tools/transformer` directory:
   
   ```bash
    cd tools/transformer
   ```

2. Then run the following command:

   ```bash
   go build
   ```

This will create a binary file called `transformer` in the `tools/transformer` directory.

<!-- INTERACTIVE page step2.md END -->

<!-- INTERACTIVE page step3.md START -->

# Learn the basic meta syntax

The Sandbox Transformer uses a special syntax to define the metadata for the course. This metadata is used to define the structure of the course and the content of each page.

Here is a breakdown of the basic meta syntax. For more information on the meta syntax, see the [full documentation](https://github.com/grafana/killercoda/blob/staging/docs/transformer.md)



## Metadata

You specify Killercoda tutorial metadata in the source file front matter as the value for the `killercoda` field.
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
```

## Directives

Directives in the source file modify how the transformer tool generates the tutorial.
You write directives in the source file with HTML comments.

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

### Exec

Exec directives tell the transform tool to make the contained fenced code block executable.

> [!NOTE]
>
> By default, the tool makes `bash` fenced code blocks executable so you don't need `<!-- INTERACTIVE exec START/STOP -->` directives for bash code blocks.
> You can override this behavior with the `<!-- INTERACTIVE copy START/STOP -->` directives which take precedence over the default behavior.

The start marker is:

```markdown
<!-- INTERACTIVE exec START -->
```

The end marker is:

```markdown
<!-- INTERACTIVE exec END -->
```

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

## Examples

The best place to see how the meta syntax works is to look at the examples in the `docs/examples` directory of the `killercoda` repository. You can find examples of how to structure your markdown files and how to use the meta syntax to define the structure of your course.

<!-- INTERACTIVE page step3.md END -->

<!-- INTERACTIVE page step4.md START -->

# Use the Sandbox Transformer to create a course

Now that you have the transformer built and you understand the basic meta syntax, you can use the Sandbox Transformer to turn a Markdown docs into a course. Lets use one of the examples in the `docs/examples` directory of the `killercoda` repository:

1. Navigate back to the root of the `killercoda` repository:

   ```bash
   cd ../..
   ```
1. Make a new directory for your new courses:

   ```bash
   mkdir new-courses
   ```
   This is where your courses for a specific topic will live. You can create multiple courses in this directory.

1. Create a new directory for your course: 

   ```bash
   mkdir new-courses/new-course-1
   ```

1. We will also create a `structure.json` file more on this later:

   ```bash
   touch new-courses/structure.json
   ```

1. Its time to run the transformer on the example course:

   ```bash
   ./tools/transformer/transformer docs/examples/complete-docs-example.md new-courses/new-course-1
   ```
   This will transform the `complete-docs-example.md` file into a course in the `new-course-1` directory.

1. Verify the course was created:

   ```bash
   ls new-courses/new-course-1
   ```
   You should see a number of files and directories created for the course.

1. Finally, add the course to the `structure.json` file:

   ```json
    {
        "items": [
        { "path": "new-course-1", "title": "New Course 1" }
        ]
    }
   ```
   This will tell Killercoda where to find the course.


<!-- INTERACTIVE page step4.md END -->


<!-- INTERACTIVE page step5.md START -->

# Testing the course

Before you open a PR to the `killercoda` repository, you should test the course to make sure it works as expected. The easiest way to do this is to run the course via your own Killercoda instance. To do this follow these steps:

1. Fork the `killercoda` repository to your own GitHub account. This will provide you with a URL to your forked repository.
   ```
   https://github.com/<USERNAME>/killercoda.git
   ```
   
1. Add the forked repository as a remote to your local repository:

   ```bash
   git remote add forked https://github.com/<USERNAME>/killercoda.git
   ```

1. Add the changes to your forked repository:

   ```bash
   git add .
   git commit -m "Add new course"
   ```

1. Push the changes to your forked repository:

  ```bash
   git push forked my-new-course
  ```

1. Create a Killercoda account: [https://killercoda.com/login](https://killercoda.com/login)

1. Then head to: [https://killercoda.com/creator/repository](https://killercoda.com/creator/repositor) and add your forked repository.

1. Once saved, you should see your course in the list of courses. Click on the course to open it.

<!-- INTERACTIVE page step5.md END -->

<!-- INTERACTIVE page finish.md START -->

# Conclusion

In this tutorial, you learned how to use the Sandbox Transformer to turn Markdown docs into a course. You learned how to build the transformer, the basic meta syntax, and how to use the transformer to create a course. You also learned how to test the course using your own Killercoda instance.

## Next steps

When you are ready, you can open a PR to the `killercoda` repository to add your course. We are excited to see what you create and to get your feedback on the Sandbox Transformer.

<!-- INTERACTIVE page finish.md END -->





