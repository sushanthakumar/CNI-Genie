# Feature 4: CNI-Genie "Default Plugin Selection"

# Background

There are many cases/user scenarios where we want pod to have ip(s) from a default network which we choose. For this case, CNI Genie provides a very useful default plugin support feature.
For using this feature, we can update the genie conf file to set the plugin of our choice as default plugin

Once this is set, we dont have to make any cni changes to the pod yaml while creating subsequent pods and the ip(s) from default plugins will be ensured
This config can be updated/modified at any point of time and will reflect in subsequent run

# Modification to genie conf file
We need to add *"default_plugin": "weave,flannel"* to the genie conf file as shown below

