package updateCustomerById_test

type expect struct {
	given *given
	when  *when
}

func NewExpect(given *given, when *when) *expect {
	return &expect{
		given: given,
		when:  when,
	}
}
