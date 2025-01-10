# Testing the course

Before you open a PR to the `killercoda`{{copy}} repository, you should test the course to make sure it works as expected. The easiest way to do this is to run the course via your own Killercoda instance. To do this follow these steps:

1. Fork the `killercoda`{{copy}} repository to your own GitHub account. This will provide you with a URL to your forked repository.

   ```
   https://github.com/<USERNAME>/killercoda.git
   ```{{copy}}

1. Add the forked repository as a remote to your local repository:

   ```bash
   git remote add forked https://github.com/<USERNAME>/killercoda.git
   ```{{exec}}

1. Add the changes to your forked repository:

   ```bash
   git add .
   git commit -m "Add new course"
   ```{{exec}}

1. Push the changes to your forked repository:

```bash
 git push forked my-new-course
```{{exec}}

1. Create a Killercoda account: [https://killercoda.com/login](https://killercoda.com/login)

1. Then head to: [https://killercoda.com/creator/repository](https://killercoda.com/creator/repositor) and add your forked repository.

1. Once saved, you should see your course in the list of courses. Click on the course to open it.
