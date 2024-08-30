---
layout: default
title: Importing a Design
abstract: Import your existing application definitions and infrastructure configuration into Meshplay.
permalink: extensions/importing-a-design
language: en
display-title: "false"
list: include
type: extensions
category: meshmap
---

# Importing a Design

You may bring your own design files or find them available through your chosen provider. Each design carries a unique identifier for reference. The designs in this repository serve in an educational capacity, facilitating learning, and also serve in an operational capacity, facilitating implementation and validation of your cloud native deployment’s adherence to a design.

**Step 1: Access the Extensions UI**

From Meshplay Extension, designs can be imported from your local filesystem or imported from a remote URL.

**Step 2: Navigate to the Designs Tab and Select your Import**

Once you have accessed the Extension's UI, navigate to the Designs tab. In this tab you can see all your designs with their "<b>Name</b>" and "<b>Date Modified</b>". From the top right of the table click on import design which opens import modal.

<a href="{{ site.baseurl }}/assets/img/meshmap/application-tab.png"><img style="border-radius: 0.5%;" alt="Import-Application" style="width:800px;height:auto;" src="{{ site.baseurl }}/assets/img/meshmap/design.png" /></a>

**Step 3: Import the Design**

You can import your design by clicking on the “Browse” button and selecting the file from your local machine or import in through URL Once you have selected the file, click on the “Import” button to import design into Meshplay Extension. When you import a design into Extensions, it will create a Meshplay Design based on definition. This Meshplay design will include all of the configurations, and other parameters defined in the File.

<a href="{{ site.baseurl }}/assets/img/meshmap/apps-modal.png"><img style="border-radius: 0.5%;" alt="Import-Application" style="width:800px;height:auto;" src="{{ site.baseurl }}/assets/img/meshmap/import-design.png" /></a>

Once the Meshplay Design has been created, you can use Meshplay Extension to manage, operate and observe your cloud native infrastructure. You can also use Meshplay Extension to deploy your Meshplay Design to your infrastructure.

