# Getting Started with the Project Setup

In this first step, we'll get your environment set up by cloning the repository, installing necessary dependencies, and starting both the frontend and backend servers. Follow the instructions below to ensure everything is set up correctly before we proceed to more advanced configurations.

## Step 1: Clone the Repository

Start by cloning the repository from GitHub to get the necessary project files. Open your terminal and run the following command:

```bash
git clone https://github.com/tomglenn/fo11y-demo.git
```{{execute}}

This command clones the `fo11y-demo` repository into a folder named `fo11y-demo` on the machine.

## Step 2: Configuring the Frontend

The frontend React application talks to our backend application, which usually runs on `localhost:1337`. However, in Killercoda the application will be running on a dynamically generated URL, so you need to update the `.env` file to point to your unique backend service URL.

Run the following command to update the `.env` file:

```bash
echo "REACT_APP_BACKEND_URL={{TRAFFIC_HOST1_1337}}" > fo11y-demo/frontend/.env
```{{execute}}

## Step 3: Starting the Frontend and Backend

We'll use Docker Compose to run both the frontend and backend applications.

Change into the `fo11y-demo` folder and then run `docker-compose`:

```bash
cd fo11y-demo
docker-compose up --build -d
```{{execute}}

*Note: This step may take quite a while, so feel free to grab yourself your favourite drink and relax.* 😎

Once both Docker containers are up and running (You can check with `docker ps`), you can check that the backend API is returning data by visiting [http://locahost:1337/featured]({{TRAFFIC_HOST1_1337}}/featured).

You should now also be able to open the frontend application at [http://locahost:3000]({{TRAFFIC_HOST1_3000}}) and see a list of featured games.