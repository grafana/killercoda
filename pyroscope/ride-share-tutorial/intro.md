# Ride share tutorial with Pyroscope

This tutorial demonstrates a basic use case of Pyroscope by profiling a “Ride Share” application.
In this example, you learn:

- How an application is instrumented with Pyroscope, including techniques for dynamically tagging functions.

- How to view the resulting profile data in Grafana using the Profiles View.

- How to integrate Pyroscope with Grafana to visualize the profile data.

## Background

In this tutorial, you will profile a simple “Ride Share” application. The application is a Python Flask app that simulates a ride-sharing service. The app has three endpoints which are found in the `server.py`{{copy}} file:

- `/bike`{{copy}}    : calls the `order_bike(search_radius)`{{copy}} function to order a bike

- `/car`{{copy}}     : calls the `order_car(search_radius)`{{copy}} function to order a car

- `/scooter`{{copy}} : calls the `order_scooter(search_radius)`{{copy}} function to order a scooter

To simulate a highly available and distributed system, the app is deployed on three distinct servers in 3 different regions:

- us-east

- eu-north

- ap-south

This is simulated by running three instances of the server in Docker containers. Each server instance is tagged with the region it represents.

![Getting started sample application](https://grafana.com/media/docs/pyroscope/ride-share-demo.gif)

In this scenario, a load generator will send mock-load to the three servers as well as their respective endpoints. This lets you see how the application performs per region and per vehicle type.

> **Tip:**
> A setup script runs in the background to install the necessary dependencies. This should take no longer than 30 seconds. Your instance will be ready to use once you `Setup complete. You may now begin the tutorial`{{copy}}.
