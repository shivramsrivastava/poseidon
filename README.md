# Introduction
The Poseidon/Firmament scheduler incubation project is to bring integration of Firmament Scheduler (OSDI paper) in Kubernetes.
At a very high level, Poseidon/Firmament scheduler augments the 
current Kubernetes scheduling capabilities by incorporating a new 
novel flow network graph based scheduling capabilities alongside the default Kubernetes Scheduler. 
Firmament models workloads and clusters as flow networks and runs min-cost flow optimizations over these networks to make scheduling decisions.

Due to the inherent rescheduling capabilities, the new scheduler enables a globally optimal scheduling environment that constantly keeps refining the workloads placements dynamically.

As we all know that as part of the Kubernetes multiple schedulers support, each new pod is typically scheduled by the default scheduler, but Kubernetes can be instructed to use another scheduler by specifying the name of another custom scheduler (Poseidon in our case) at the time of pod deployment. In this case, the default scheduler will ignore that Pod and allow Poseidon scheduler to schedule the Pod on a relevant node.

# Key Advantages

* Flow graph scheduling provides the following 
  * Support for high-volume workloads placement 
  * Complex rule constraints 
  * Globally optimal scheduling
  * Extremely high scalability. 
  
  **NOTE:** Additionally, it is also very important to highlight that Firmament scales much better than default scheduler as the number of nodes increase in a cluster.

An argument can be made from functional parity standpoint between two scheduling algorithms at this point of time. Currently, Firmament functionally is not at par with the default scheduler, especially for affinity/anti-affinity functionality as well as various other functional features currently supported in the default scheduling algorithm. It would definitely take some time to bridge the functionality gap between two scheduling algorithms. But our argument is that due to the exceptional throughput superiority of Firmament scheduler, certain use cases may not even require additional functional features such as advanced affinity/anti-affinity etc.


For more details about the design of this project see the [design document](https://docs.google.com/document/d/1VNoaw1GoRK-yop_Oqzn7wZhxMxvN3pdNjuaICjXLarA/edit?usp=sharing) doc.


![image](https://drive.google.com/open?id=1Yy-WI5cOIXNR5NANNsDyWShhrEjfqC1I)

# Installation
  In cluster installation of poseidon, this method is to run posiedon and its components in 
  kubernetes cluster.
  
  The following are the componentes required.
   * [Running Kubernetes cluster](https://kubernetes.io/docs/setup/).
   * [Firmament scheduler](https://github.com/Huawei-PaaS/firmament) deployment
   * Poseidon deployment
   * Heapster with the poseidon sink.
   
  Please start [here]().
  
  
  
# Development
  For developers please refer [here]()


# Roadmap
  * Affinity and Anti-Affinity implementation
  * CI and CD integration
  * Documentation improvements
  
