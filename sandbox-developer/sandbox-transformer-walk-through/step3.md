# Use the Sandbox Transformer to create a course

Now that you have the transformer built and you understand the basic meta syntax, you can use the Sandbox Transformer to turn Markdown docs into a course. Lets use one of the examples in the `docs/examples`{{copy}} directory of the `killercoda`{{copy}} repository:

1. Navigate out of the `killercoda`{{copy}} repository:

   ```bash
   cd ../
   ```{{exec}}

1. Make a new directory for your new courses:

   ```bash
   mkdir killercoda/new-courses
   ```{{exec}}

   This is where your courses for a specific topic will live. You can create multiple courses in this directory.

1. Create a new directory for your course:

   ```bash
   mkdir killercoda/new-courses/new-course-1
   ```{{exec}}

1. We will also create a `structure.json`{{copy}} file more on this later:

   ```bash
   touch killercoda/new-courses/structure.json
   ```{{exec}}

1. Its time to run the transformer on the example course:

   ```bash
   ./transformer killercoda/docs/examples/complete-docs-example.md killercoda/new-courses/new-course-1
   ```{{exec}}

   This will transform the `complete-docs-example.md`{{copy}} file into a course in the `new-course-1`{{copy}} directory.

1. Verify the course was created:

   ```bash
   ls killercoda/new-courses/new-course-1
   ```{{exec}}

   You should see a number of files and directories created for the course.

1. Finally, add the course to the `structure.json`{{copy}} file:

   ```json
    {
     "items": [
     { "path": "new-course-1", "title": "New Course 1" }
     ]
    }
   ```{{copy}}

   This will tell Killercoda where to find the course. This can be done via the inbuilt code editor in the Killercoda UI. Or you can use nano:

   ```bash
    nano killercoda/new-courses/structure.json
   ```{{exec}}
