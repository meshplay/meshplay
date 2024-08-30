---
layout: default
title: Importing existing applications
permalink: extensions/importing-an-application
language: en
abstract: Import your existing applications into Meshplay.
display-title: "false"
list: include
type: extensions
category: meshmap
---

## Import existing applications into Meshplay

Importing apps into Meshplay Extension as Meshplay Apps is a powerful feature that enables users to manage, operate, and observe their cloud native infrastructure more effectively. With this feature, users can easily import their app manifest and store it in the database.

**Step 1: Access the Extensions UI**


Note: You can import three types of application: Docker Compose, Helm Chart, Kubernetes Manifest.


**Step 2: Navigate to the Applications Tab and Select your Import**


Once you have accessed the Extension's UI, navigate to the Applications tab. In this tab you can see all your application with their "<b>Name</b>" and "<b>Date Modified</b>". From the top right of the table click on import application which opens import modal.

<a href="{{ site.baseurl }}/assets/img/meshmap/application-tab.png"><img style="border-radius: 0.5%;" alt="Import-Application" style="width:800px;height:auto;" src="{{ site.baseurl }}/assets/img/meshmap/application-tab.png" /></a>


**Step 3: Import the Application**

You can import your application by select File Type from the options and clicking on the “Browse” button and selecting the file from your local machine or import in through URL Once you have selected the file, click on the “Import” button to import app into Meshplay Extension. When you import a app into Extensions, it will create a Meshplay App based on definition. This Meshplay App will include all of the services, ports, and other parameters defined in the File.

<a href="{{ site.baseurl }}/assets/img/meshmap/apps-modal.png"><img style="border-radius: 0.5%;" alt="Import-Application" style="width:800px;height:auto;" src="{{ site.baseurl }}/assets/img/meshmap/apps-modal.png" /></a>

Once the Meshplay Application has been created, you can use Meshplay Extension to manage, operate and observe your cloud native infrastructure. You can also use Meshplay Extension to deploy your Meshplay Application to your infrastructure.

