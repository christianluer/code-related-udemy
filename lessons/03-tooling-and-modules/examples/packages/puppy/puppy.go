// Package puppy is a tiny package used to demonstrate imports and visibility.
package puppy

import "github.com/christianluer/code-related-udemy/lessons/03-tooling-and-modules/examples/packages/dog"

// Bark returns a puppy bark.
func Bark() string { return "Woof! " + secret() }

// Barks returns several barks.
func Barks() string { return "Woof! Woof! Woof!" }

// BigBark uses package dog, which makes dog a dependency of puppy.
func BigBark() string { return dog.WhenGrownUp(Bark()) }

// From10 mimics the course's versioned function (tag v1.0.0, v1.1.0...).
func From10() string { return "I'm from version 1.0.0" }

// secret is unexported: lowercase means not visible outside package puppy.
func secret() string { return "(wags tail)" }
