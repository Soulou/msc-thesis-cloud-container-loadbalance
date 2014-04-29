## Papers and summary by category

### Background

#### Virtualisation

> __Tags__: background, operating system virtualisation

__[1] Paul Barham, Boris D., Keir F., Steven H., Tim H., Alex H., Rolf N., Ian P., Andrew W.<br>
  Xen and the Art of Virtualisation<br>
  19th ACM Symposium on Operating Systems Principles - 2003__

This paper details how operating system and hardware virtualisation is
basically working in order to isolate process and slice of resources from each
other.


#### Live migration
 
> __Tags__: background, live migration

__[2] Christopher Clark, Keir F., Steven H., Jacob G. H., Eric J., Christian L., Ian P., Andrew W.<br>
  Live migration of virtual machines<br>
  Proc. of the 2nd Symposium on Networked Systems Design and Implementation (NSDI’05), 2005.__

The publications details the workflow for migrating a virtual machine from on host
to another and look at the impact on working services as web server of video game server
in order to measure the migration overhead. The process of copying memory pages then pausing
the source server and enabling the destination server and redirecting the network traffic
is explained.

#### Bin packing

> __Tags__: background, online bin packing

__[3] G. Gambosi, A. Postiglione, and M. Talamo<br>
  Algorithms for the relaxed online bin-packing model<br>
  SIAM J. Comput. Issue 5, vol. 30, 2000.__

The relaxed online bin packing algorithm define a way to achieve online bin
packing with extra move of items when packing a new item. The number of
additional move is bounded to 3 for inserting a new element and 7 when an
element has to move.  Extrapolated to virtual machines migrations, it's
interesting because it allows a limited amount of VM migration when adding a
new VM, it is not strict as a simple online bin packing algorithm.

> __Tags__: background, online bin packing

__[4] Joseph Wun-Tat Chan Tak-Wah Lam Prudence W.H. Wong
  Dynamic Bin Packing of Unit Fractions Items
  Theoretical Computer Science issue 8, vol. 409, 2008__

Dynamic online bin packing differenciate itself from standard bin packing by
allowing items to be removed from bins. Static online bin packing doesn't allow
these items changes, once an item has been placed it doesn't move anymore. The
publication compares the dynamic bin packing to offline bin packing in order to
calculate its 'N-competitive' value.

Use of the publication: dynamic online bin packing is quoted by different other
paper as an interesting algorithm, that is why it may be interesting to add in as
a background paper.

### Motivation

> __Tags__: virtualization, resiliency

__[5] Salapura, V.
  Cloud computing: Virtualization and resiliency for data center computing
  Computer Design (ICCD), 2012 IEEE 30th International Conference__

Small publication explaining the importance of virtualisation and why companies
should virtualize their infrastructure in order to setup scalable, highly
available services based on cloud infrastructure.  It also focuses on how
disaster recovery scenarios are simplified and cheaper thanks to it. These reasons
are basic arguments for using cloud infrastructure and virtual machines to run
workloads.

- - -

> __Tags__: reassignment necessity, live migration, energy saving

__[6] Thomas, S., Alexander, S.
  Decision support for virtual machine reassignments in enterprise data centers
  2010 - Network Operations and Management Symposium__

Thomas S. starts from the statement that energy represents up to 50% of
operating costs in infrastructure, that's why there is a need to optimize it.
By using reassignment with live migration, they are looking at hosts
consolidation in order to minimize the number of active (powered) physical
hosts.

Different numbers of way to execute the consolidation are listed, first using
__first fit algorithm__ in a reactive way, when some metrics
(CPU/memory/IO/bandwidth usage) are over a precise level over a certain period.
But the result is suboptimized because only 1 parameter is taken into account.
Then vector __bin packing__ is introduced and the fact that it is NP-hard, so
difficult to apply on large set of instances without good heuristics.  Thus, an
approached by defining a model of VM consumption by time range of 5 minutes for
different resources and using the complementarity to optimize the consolidation
is defined. For example high-CPU low-memory with low-cpu high-memory processes.
(30-35% of PM may be saved using this complementarity)

This is to introduce re-assignment based on virtual servers migrations
considering the migration overhead which can be quite important (according to
size of VM) and which may slow down other VMs on the host.

Experiment: "From a professional data center we obtained data describing
one-hour maxima for CPU demand of hundreds of VMs over four months."

This paper is not focusing which VM has to be moved or on which host it should
be done, but it shows the necessity of re-assignment by analyzing the CPU load
on a hourly basis.

### Algorithmic study - methods

> __Tags__: reassignment, online bin packing

__[7] Weijia Song, Zhen Xiao, Qi Chen, Haipeng Luo<br>
  Adaptive Resource Provisioning for the Cloud Using Online Bin Packing<br>
  2013 - IEEE Transactions on Computers__

Introduce topic with Amazon EC2 80k launched VMs per day and half milion PM
estimated and that one of main cost is electricity saving, so optimization is
required, so using live migration is great to consolidate.

Virtual machine allocation using bin packing. Focus on unidimensional vector
bin packing. Extend it to variable size vector been packing, however, they are
not considering all the dimensions but only the most important in a vector.  To
predict the load, they use a variant of the exponentially weighted moving
average (EWMA) load predictor.

In this case, they assume from [2], that migrating VMs doesn't create important
overhead to affect the other VMs, so they don't consider it. They use a
scheduler which is invoked periodically to predict the use of VMs according to
former saved statistics (CPU/Memory/NetworkIO) and the output is a list of
migrations to execute.

Issue with offline bin packing: it can create a lot of change in the VMs layout
when executing the algorithm because it does not consider the current layout,
that is why online bin packing is much more realistic. Strict online bin
packing doesn't allow to move existing items, it's too restrictive. Relaxed bin
packing only allow a constant number of items to be repack [3], so it doesn't
fit either. Dynamic online bin packing [4] is considered, but discarded. So the
designed solution is called __Variable Item Size Bin Packing__

VISBP:
* Insert similar to relaxed online bin packing (no knowledge of future when
  item is packed) [2]
* Insert and Change functions (no remove like dynamic bin packing have)
* During an operation a small amount of additional moves is allowed

The algorithm is extended to multi-dimensional bin-packing to consider
different types of resources. However, they are flattening it using the biggest
dimension of each item. So when there is a big difference between the different
dimensions or when the number of dimensions is high, this algorithm is not
working well.

Simulation:
Trace from university, all the services DNS, mail, syslog, any kind of service.
(500 datasets, 60 PM, 1200VM) The comparaison is done with offline bin packing,
black box gray box and VectorDot. The results are based on the number of active VM,
and de facto, the number of active PMs. The fewer, the less electricity is consumed.

Experiments

Small scale experiment:

* 3 VMs on 1 PM. Show Xen schduler limiting at 80% the CPU load when the 3 VMs are doing
high CPU usage
* 3 PMs - 5 VMs. Unidimensional (CPU) VISBP
* 2 PMs - 5 VMs. Multidimensional (CPU-NetworkIO) VISBP

Algorithms comparaison:

* 10PMs (10GB, 4CPUs -no hyperthreading), 80VMs (2GB, 1vCPU). 
* Benchmarking with TCP-W
* Server: online bookstore, each VM is a copy, workload: consumers buying books
* Init: 4PMs with 15VMs, 4PMs with 7-6-4-3VMs, scheduling every 10 minutes

Results: Less active PMs with VISBP and more PMs over 25% of CPU usage after consolidation

### Experimental study

- - -

> __Tags__ : offline vector bin packing, heterogeneous environment

__[8] Mark, S., Frédéric, V., Henri, C. <br>
   Virtual Machine Resource Allocation for Service Hosting on Heterogeneous Distributed Platforms<br>
   2012 - Parallel & Distributed Processing Symposium__

In contrast with the previous publication, the paper deals with offline bin
packing based allocation, But it adds another constraint: the heterogeneity of
the platform. Three different kind of algorithm are compared. Exact solution
based on the linear programming definition of the resource allocation on
heterogeneous host. The main issue is that as the problem is NP-hard, finding
the exact solution takes a lot of time and this operation should be quick. Then
greedy algorithms are studied, they are fast to execute, but the results are
not really good. Finally, the main topic of this publication, the vector
packing algorithms; they lead to the best yield. (The yield represents the
relative performance of a serve, the closer to 1 is the yield, the better the
performance)

The algorithms for vector packings are called METAVP and METAHVP for its
heterogeneous version. Why meta, because the gather several times the same
algorithm using different initial conditions. These conditions are based on the
way to sort the PMs (bins) and the VMs (items) before starting the packing. The
best result is used as a result for the META-H-VP result. Later the
META-H-VPLIGHT algorithms are defined. They are similar to the previous ones
except that the least efficient sorting types have been removed. (with 512PMs
and 2000VMs, the simulation using META-H-VPLIGHT takes 10 times less time than
simple META-H-VP)

Results: VP algorithms are more efficient that greedy/linear programming.
On heterogeneous environment they do much better that the others. Bruteforce
approach take dereasonable time, but a light version can be define, dividing
the runtime by 10

Error in CPU estimation: In practice, it is difficult to have a right
estimation of the future CPU consumption and a bad estimation can result in a
really bad placement of services on the PMs. However an approach is proposed to
give a default value to these tasks and the obtained result is better than
zero-knowledge scheduling (from "Scheduling in the Dark")

- - -

> __Tags__: heterogeneous environment, provisioning, google trace

__[9] Qi Zhang, Mohamed Faten Zhani, Raouf Boutaba, Joseph L. Hellerstein
  HARMONY: Dynamic Heterogeneity−Aware Resource Provisioning in the Cloud
  2013 - IEEE 33rd International Conference on Distributed Computing Systems__

HARMONY → Heterogeneity Aware Resource MONitoring and management sYstem
Publication done by the University of Waterloo (CA) and Google. Based on real
values from Google datacenters. The concept of heterogeneity has been split
between Machine Heterogeneity, which consists in having physical servers with
different specifications and Workload Heterogeneity, having tasks having
completely different requirements. (Duration, Load, IO, etc.)

The study of Google's workload shows that they are working mostly on small
tasks with low memory consumption and low CPU usage but very numerous.

As, at Google, the granularity of the task is too small, they are allocating
containers which will handle tasks. Their problem of optimization is to avoid
over-allocating and wasting resources.

" Finally, traditional bin−packing heuristics (e.g., First−Fit) do not apply
directly to CBS as they do not consider machine switching and container
reassignment costs. " (CBS: Container Based Scheduler)

> __Tags__: network flow, theory, resource allocation

__[10] Kimish Patel, Murali Annavaram, Massoud Pedram
  NFRA: Generalized Network Flow Based Resource Allocation for Hosting Centers
  2013 - Transactions on Computer__

NFRA: Network Flow Resource Allocation


- - -

> __Tags__: network flow, vm allocation

__[11] Fangzhe Chang, Jennifer Ren, Ramesh Viswanathan
  Optimal Resource Allocation in Clouds
  2010 - IEEE 3rd International Conference on Cloud Computing__

TODO

- - -

> __Tags__: reassignment, group of vms

__[12] Hien Nguyen Van, Frédéric Dang Tran, Jean-Marc Menaud
  SLA-aware Virtual Resource Management for Cloud Infrastructures
  2009 - IEEE Ninth International Conference on Computer and Information Technology__

Work around the virtual machine allocation. They separe two components: first
the Local Decision Module (LDM) and the Global Decision Module (GDM). A LDM is
responsible of an Application Environment (AE) which gathers all the resources
needed by a given application (i.e. multi-tier applications). For example a LDM
can manage 3 VM which are instances of a service and the related database. The
LDM is in charge of allocating or deprovisionning virtual machines for a given
application. The GDM is responsible for managing the request from the LDMs and
doing the best choices for VM allocation for each application.

The main goal is virtual machines consolidation, reducing the number of required
physical machines without overloading any of them and respecting the SLA of
each hosted application.

The mathematical approach is done by solving an optimization problem. It defines
two Constraint Satisfactions Problems. The first is for the VM allocation and the
second for the VM packing, the consolidation.

> __Tags__: reassignment, reactive, cpu-based

__[13] Mauro, A., Sara, C., Michele, C., Michese, M.
  Dynamic load management of virtual machines in a cloud architecture
  2010 - Social-Informatics and Telecommunications Engineering__

This publication defines a model to manage virtual machines migrations. 1.
Select overloaded hosts (host monitoring), 2. Select guests to migrate (VM
monitoring), 3. Select new host, 4. Migrate the VM. It doesn't use any
innovative algorithm for the host selection. They choose sort the nodes by
load, and migrate up to one VM to each of these nodes, from the more available
to the more loaded. However they are only considering the CPU load. It is the
most important metric to keep guests alive and fast, but memory, network
bandwidth and I/O may influence these results.

- - -

> __Tags__: reassignment

__[14] Ramon, L., Vinicius W.C., M., Thiago, F., N., Vitor A.A., S.,
  Heuristics and matheuristics for a real-life machine reassignment problem
  2013 - International Transactions in Operational Research__

Formal answer to the problem proposed by Google during the ROADEF/EURO
Challenge. The solution is based on a linear integer programming (IP)
formulation and Iterated Local Search (ILS) heuristics. They solve the
reassignment on an infrastructure of 50,000 tasks based on 5,000 hosts. The
approach is interesting even if it is different from bin packing (quoted in the
conclusion)

- - -

> __Tags__: reassignment, live migration, overhead

__[15] Thomas, S., Andeas, W.
  Virtual Machine Re-Assignment Considering Migration Overhead
  2012 - Network Operations and Management Symposium__

Short publication (may be taken as an extension of the previous publication)
which stats at the observation that in publications relative to virtual machine
assignment do not take care of the migration overhead in their model.  They try
to define a model which contains constraints related to the timing or number of
migrations.

### Surveys - Literature information

> __Tags__: survey, resource allocation

__[16] Swapnil M Parikh
  A Survey on Cloud Computing Resource Allocation Techniques
  2013 - Nirma University International Conference on Engineering__

This publication gathers recent papers on resource allocation in the scope of cloud computing.
It introduces to some other approaches such as "game-theoretic method for fair
resource allocation in cloud computing" for example.
