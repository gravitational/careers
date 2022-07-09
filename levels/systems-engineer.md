
# Engineering Levels

The matrix below explains the requirements for engineering levels. Each level
has to match at least one of the requirements for the level and the level before it.

Not all skills go to level 6. This is determined by the focus of Teleport.

For example, security skills are more relevant, as we are a security oriented
company, and we are interested in R&D level 6 engineers, however OS/tooling is
not the primary focus, so basic levels are satisfactory.

As the level advances past 4, the focus shifts from technical skills to bringing
products to market, as it's more relevant to Teleport as a for-profit startup.

Skills are a mix of technical and communication abilities. Some references to
security are based on the coding challenge.


# Levels

## L1

### Data structures/Algorithms
Demonstrates good understanding of basic data structures - hash tables, linked
lists, trees. Can reason about algorithm complexity. Applies relevant data
structures in day to day activity, can implement a production quality. It might
be not the most efficient or secure, but correct.

### Product Output
Creates a design document based on well-defined scoped requirements and implements it.

### Security
Applies basic security principles to program design. For example, can set up HTTPS
and password based auth.

### Communication/Writing
Reports progress on a regular basis as required by the team's operational
requirements. Actively solicits feedback. Participates in interview panels.

### Networking
Understands and reasons about networking concepts. Understands and can write
production quality web servers. Understands common networking issues and
troubleshooting techniques.

### Systems Engineering
Understands the usage of POSIX and other APIs for Linux systems. Understands
synchronization primitives and their application, including reasoning about
deadlocks and data races. Can write basic system-level code using the different
types of memory and allocation. Understands inter process communication and can
build systems leveraging it. Can implement data race and deadlock free code
using basic production guidelines - using synchronization primitives and
properly sharing state between components of the system.

### OS/Tooling
Understands the usage of compilers, interpreters, build tools at the organization.

## L2

### Data structures/Algorithms
Same as L1.

### Product Output
Same as L1.

### Security
Applies industry best practices and security guidelines, like setting up strong
TLS, can pick strong authentication and authorization mechanisms.

### Communication/Writing
Provides constructive review on peers' code and design. Helps new team members
during their first weeks.

### Networking
Has more granular understanding of the networking design, for example can reason
about using GRPC vs HTTPS-JSON and their networking and scalability trade-offs.
Can do the same for UDP vs TCP and lower level protocols.

### Systems Engineering
Can reason about performance implications and risks of using synchronization
primitives, understands granularity of locking, can debug and troubleshoot
memory and synchronization issues.

### OS/Tooling
Understands tools and compilers as an advanced user - for example, can refactor
Makefiles to make them more efficient and parallel execution friendly or select
the optimal set of compiler flags for Linux and MacOS.

## L3

### Data structures/Algorithms
Uses advanced data structures and algorithms, such as consensus, gossip
protocols to implement key product features.

### Product Output
Collaborates with the team to scope requirements, based on good understanding
of existing longer term product vision and estimates of the system design of a
feature of a product.

### Security
Understands security aspects of protocols on a deeper level, for example can
explain differences between TLS 1.2 and 1.3 and security implications.
Understands common attack vectors for server side or client side applications
and can build secure systems that will pass quality security audit that will
uncover no critical system design errors.

### Communication/Writing
Supports less experienced peers' technical skills, answering questions and being
a resource. Documents and improves team practices.

### Networking
Understands scalability and resilience aspects of practical implementation of
systems leveraging networking. Can write fast, scalable servers and troubleshoot
common networking issues such as connection lifecycle, pooling, backpressure and
other aspects of the application design.

### Systems Engineering
Not only can write data-race and deadlock free code, but implements safe and
concurrent and/or parallel systems using minimum amount of shared state, granular
locking - systems that are easy to read, extend and troubleshoot. Senior
engineers are not expected to find any faults or provide any further suggestions
in the systems design document submitted by the Level 3 engineer.

### OS/Tooling
Same as L2.

## L4

### Data structures/Algorithms
Leads implementation of products and standalone systems using advanced data
structures and algorithms, such as upgrade/install distributed framework.

### Product Output
Leads the implementation of the isolated feature/improvement that measurably and
significantly impacts business outcomes from gathering requirements to getting to
the market stage.

### Security
Can apply production quality novel cryptographic to build security systems. For
example, can implement strong security support for the system with eBPF from
scratch.

### Communication/Writing
Writes technical articles/blog posts, delivers tech and lightning talks
representing the company's technical vision.

### Networking
Same as L3.

### Systems Engineering
Can implement production grade systems leveraging advanced low-level and novel
components of the Linux design like BPF, control groups.

### OS/Tooling
Same as L3.

## L5

### Data structures/Algorithms
Implements customer facing features and systems using advanced data structures
and algorithms.

### Product Output
Leads the implementation of a new product line or significant part of the
product to deliver it to the market in collaboration with all other teams.

### Security
Writes technical articles on security aspects of the system, implements
significant security product innovations in the area delivered to customers.

### Communication/Writing
Writes advanced technical articles/blog posts, gaining significant industry
traction or delivers technical talks on major conferences representing the
company's vision.

### Networking
Same as L4.

### Systems Engineering
Applies system level design to deliver new products or significant new
components of an existing product to the market.

### OS/Tooling
Same as L4.

## L6

### Data structures/Algorithms
Designs new data structures and algorithms solving relevant business problems
and creating competitive advantage for the company.

### Product Output
Same as L5.

### Security
Researches and designs new security systems, cryptography and protocols.

### Communication/Writing
Produces peer-reviewed research papers or patent applications.

### Networking
Same as L5.

### Systems Engineering
Same as L5.

### OS/Tooling
Same as L5.
