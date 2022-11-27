# medium

## Content
- [Solid Design Principle](#Solid-Design-Principle)

# Solid-Design-Principle
- The acronym SOLID was derived by Robert C.Martin also known as Uncle Bob. It defines the principle used in object-oriented design. These principles serve as a guide for software developers to build applications that are extensible and maintainable. Although Golang is partly object-oriented, as such it does not implement principles like inheritance, these SOLID principles can still be utilized with Go to build maintainable applications.

- The SOLID means:

- S: Single Responsibility Principle (SRP)
- Single Responsibility Principle (SRP): This states that a type should have one primary responsibility, as a result, it should only have one reason to change which should be the primary responsibility. For instance, let’s assume you have a chef, his or her responsibility is to prepare meals, and delicacies for you. For no reason should the chef start styling your hair or decide to start doing your laundry. This is exactly what the SRP is trying to explain. This SRP principle extends to another principle called SOC (separation of concerns). This simply means different problems or different concerns/functions that a system solves should reside in a different construct or package. To further understand these principles I have a written a simple example of a Journal that shows how the SRP is implemented

- O: Open-Closed Principle(OCP)

- L: Liskov Substitution Principle(LSP)

- I:Interface Segregation Principle(ISP)

- D: Dependency Inversion Principle(DIP)

