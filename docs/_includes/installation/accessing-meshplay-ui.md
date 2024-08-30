{% if page.display-title != "false" %}
## Accessing Meshplay UI
{% endif %}

After successfully deploying Meshplay, you can access Meshplay's web-based user interface. Your default browser will be automatically opened and directed to Meshplay UI (default location is [http://localhost:9081](http://localhost:9081)).

You can use the following command to open Meshplay UI in your default browser:

{% capture code_content %} meshplayctl system dashboard {% endcapture %}
{% include code.html code=code_content %}

If you have installed Meshplay on Kubernetes or a remote host, you can access Meshplay UI by exposing it as a Kubernetes service or by port forwarding to Meshplay UI.

{% capture code_content %} meshplayctl system dashboard --port-forward {% endcapture %}
{% include code.html code=code_content %}

Depending upon how you have networking configured in Kubernetes, alternatively, you can use kubectl to port forward to Meshplay UI.

{% capture code_content %}kubectl port-forward svc/meshplay 9081:9081 --namespace meshplay{% endcapture %}
{% include code.html code=code_content %}

<details>
<summary>Customizing Meshplay Provider Callback URL</summary>

Customize your Meshplay Provider Callback URL. Meshplay Server supports customizing authentication flow callback URL, which can be configured in the following way:

{% capture code_content %}$ MESHPLAY_SERVER_CALLBACK_URL=https://custom-host meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

<br />
Meshplay should now be running in your Kubernetes cluster and Meshplay UI should be accessible at the `EXTERNAL IP` of `meshplay` service.

</details>

Production deployments are recommended to access Meshplay UI by setting up a reverse proxy or by using a LoadBalancer.

Log into the [Provider](/extensibility/providers) of your choice.
