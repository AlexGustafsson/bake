# Bake

> A cross-platform language and tool for building things - a better Make

⚠️ Bake is currently being actively developed. Until it reaches v1.0.0 breaking changes may occur in minor versions.

Bake is a new language and toolset for building things. Think of it like Make, with more tools to empower you to easily configure pragmatic builds of smaller projects. Make is a great tool which works wonders for smaller projects. It has a great, (mostly) readable syntax which is largely natural to work with. It does, however, suffer from some shortages. For example, it does not provide more advanced scripting features for more complex builds or configurations. Furthermore, it's not easily split into reusable models. Other common and arguably more complete tools such as Gradle, Mason etc. are not as simple and playful as a simple Makefile. The ambition of Bake is to become a better Make, building on the ideas, syntax and features provided by Make.

Bake has a couple of goals:

1. Provide a familiar syntax for Make users, but enable an even more general use with further scripting capabilities
2. Provide a uniform cross-platform experience without having to install Bake itself a la Gradle
3. Enable users to create libraries for common actions, easily imported a la Go
4. Be as fast as, or faster than Make

Bake also has a non-goal:

1. Be compatible with Make
