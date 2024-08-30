---
layout: default
title: Database
permalink: concepts/architecture/database
type: components
redirect_from: architecture/database
abstract: "Meshplay offers support for internal caching with the help of file databases. This has been implemented with several libraries that supports different kinds of data formats."
language: en
list: include
---

## What are the Meshplay Databases?

Meshplay Databases function as repositories for [MeshSync](/concepts/architecture/meshsync), user preferences and system settings. Both databases are considered ephemeral and should be treated as caches. Data retention is tethered to the lifetime of their Meshplay Server instance. [Remote Providers](/extensibility/providers) may offer long-term data persistence. Meshplay's APIs offer mechanisms for clients, like [`meshplayctl`](/reference/meshplayctl) and Meshplay UI to retrieve data.

See the figure below for additional details of the data formats supported and type of data stored.

[![Architecture Diagram]({{ site.baseurl }}/assets/img/architecture/meshplay-database.svg)]({{ site.baseurl }}/assets/img/architecture/meshplay-database.svg)

### Components

Meshplay Database has several kinds of database implementations to support various usecases. They are listed below:
{% assign sorted = site.adapters | sort: "project_status" | reverse %}

| Component      | Library                               |
| :------------- | :------------------------------------ |
| Bitcask        | git.mills.io/prologic/bitcask         |
| SQLite         | gorm.io/gorm, gorm.io/driver/sqlite   |
