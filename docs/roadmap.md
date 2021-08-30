# Roadmap

## v0.1.0 - Proof of Concept

> We're here

The first usable release of bake. Makes it possible to evaluate the tool and its future.

## v0.2.0 - Robustness

This milestone will take care of technical debt left from the proof of concept phase of the project. It will include most of the planned features and ensure that they work in a robust manner.

## v0.3.0 - Imports

This milestone will implement the import functionality, with a versioned module / package system.

## v0.4.0 - Usability

This milestone will improve the overall usability of bake, with improved error messages as well as quick fixes. It will also provide language server improvements such as documentation. This milestone also marks the start of the official formatter. It will also ensure that the CLI is ergonomic.

## v0.5.0 - Dependency Tree

This milestone will implement the dependency tree to resolve in what order to build rules. It will also make it possible to perform watch builds - if a dependency change then bake may react and rebuild the necessary parts.

## v1.0.0 - Stable

_Note: There may be more releases up until v1.0.0 than listed here._

This milestone will mark the first stable release of bake and its related tools. Bake will be able to be adopted in projects. Projects may rely on their builds working today and in the future, with versioned imports etc.
