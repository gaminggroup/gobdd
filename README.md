Usage:

```go
func TestMySuites(t *testing.T) {
	MySuite()
	MySuite2()
	RunSuites(t)
}

func MySuite() {
	Suite(func() {
		Feature("My Feature 1", func() {
			Scenario("My Scenario 1.1", func() {
				Given("the system is in a state", func() {
					When("I perform an action", func() {
						Then("I expect some result", func() {
							Expect(1, Not(ToEqual(2)))
							Expect(1, ToEqual(1))
						})
					})
					When("I perform another action", func() {
						AndWhen("this action has more setup", func() {
							Then("I want to get this result", func() {
								Expect(2, ToEqual(2))
							})
							AndThen("I want to get this result", func() {
								Expect(4, Not(ToEqual("4")))
								Expect(nil, Not(ToBeNil()))
							})
						})
					})
				})
			})
		})
	})
}

func MySuite2() {
	Suite(func() {
		Feature("My feature 2", func() {
			Scenario("My Scenario 2.1", func() {
				Given("the system is in a state", func() {
					When("I perform an action", func() {
						Then("I expect some result", func() {
							Expect(1, Not(ToEqual(1)))
						})
					})
				})
			})
			Scenario("My Scenario 2.2", func() {
				Given("the system is in another state", func() {
					When("I perform an action", func() {
						Then("I expect some result", func() {
							Expect(1, Not(ToEqual(1)))
						})
					})
				})
			})
		})
	})
}

```