# entity

This is the innermost layer of the Clean Architechture.

Entities encapsulate Enterprise wide business rules. An entity can be an object
with methods, or it can be a set of data structures and functions. It doesn’t
matter so long as the entities could be used by many different applications in
the enterprise.

Entities must not depend on anything else in the project. They can depend on
outside packages, or even on the Enterprise packages that are of a very general
nature (logging, UUID, encoding, formatting etc).
