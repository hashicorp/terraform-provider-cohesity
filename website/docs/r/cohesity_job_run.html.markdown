---
layout: "cohesity"
page_title: "Cohesity: cohesity_job_run"
description: |-
  Run a Protection Job on Cohesity Cluster
---

# cohesity\_job\_run

Run a Protection Job on Cohesity Cluster

## Example Usage
```
provider "cohesity" {
        cluster_vip = "10.2.45.143"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}


resource "cohesity_job_run" "protect_vcenter" {
	name = "Job1"
	state = "start"
	timestamp= "13:00"
}
```

### Argument Reference
The following arguments are supported:

- name - (Required, string) The name of the protection job
- run_type - (Required, string) Specifies the type of backup. The default value is **Regular** The supported types include **Full**,  **Regular**, **Log** and **System**
- state - (Optional, string) Specifies whether to start or stop a protection job run. The default value is **start**. The supported values are **start** and **stop**
- operation_timeout - (Optional, int) Specifies the time to wait in minutes for the protection job run to complete the run or stop the run. The default value is **120**
- timestamp - (Required, string) Specifies the current timestamp to trigger starting or stopping a job run. The format is HH:MM


#### Attributes Reference
The following attributes are exported:

- id - ID of the VMware protection job