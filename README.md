# Trigger
<i>"Event-based automation for modern infrastructure."</i>

Trigger is a lightweight, open-source tool that allows users to define event-driven workflows in Kubernetes or Linux/cloud-native environments. Think of it as a simplified, focused version of tools like Argo Events or Tekton Triggers — but more flexible, easier to configure, and designed for sysadmins/SREs/devs who want to automate infra-related tasks based on events. You can also think about Trigger as "cool version of crontab" but for cloud-native/k8s environments

## Problems it solves:
1. Event-based Automation in K8s is Complex
Current tools like Argo Events or Tekton can be overkill for simple tasks. Users want to run a script when a pod crashes, a PVC fills up, or a new CRD is created — but writing controllers or webhooks is too heavy.

2. Bridging Infra & Dev Workflows. DevOps teams often want to trigger CI/CD or infra actions from system events (e.g. disk usage threshold, service restart, node tainting). Trigger becomes the “glue” between event sources and actions.

3. Observability + Reaction It's one thing to detect something with Prometheus or Falco. It's another to automate a reaction — restart, notify, scale, apply a patch, etc.

## Core Features
Feature	Description
📦 CRD-based config (K8s)	- Define triggers declaratively using Kubernetes Custom Resources

⚡ Linux support (CLI) -	Use it as a daemon or CLI tool on Linux hosts, outside of K8s

📥 Event sources - Pod crash, Node taint, PVC full, kube audit logs, Falco alerts, etc.

🎯 Actions (targets) -	Run job, restart pod, send webhook, Slack/Teams message, shell command

🔄 Retry/backoff policies -	Built-in support for retries, delays, circuit breakers

🧪 Dry run mode	- Test workflows without executing actions

🔐 Secure execution	- RBAC-aware, safe sandboxing of shell commands

## Archtecture 
```
+--------------------------+
|   Event Source           |
|  (K8s, Audit, Falco, etc)|
+-----------+--------------+
            |
            v
+--------------------------+
|    Trigger Engine        |
|  - Match Rules           |
|  - Evaluate Conditions   |
|  - Retry Policies        |
+-----------+--------------+
            |
            v
+--------------------------+
|     Action Handler       |
|  (Job, Webhook, Notify,  |
|   Command, etc)          |
+--------------------------+
```
