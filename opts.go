package instrumentedsql

// Opt is a functional option type for the wrapped driver
type Opt func(*wrappedDriver)

// WithLogger sets the logger of the wrapped driver to the provided logger
func WithLogger(l Logger) Opt {
	return func(w *wrappedDriver) {
		w.Logger = l
	}
}

// WithTracer sets the tracer of the wrapped driver to the provided tracer
func WithTracer(t Tracer) Opt {
	return func(w *wrappedDriver) {
		w.Tracer = t
	}
}

// DisableArgsLabel will prevent the spans from submitting
// any of the argument data with the spans in case it is sensitive data
func DisableArgsLabel(t Tracer) Opt {
	return func(w *wrappedDriver) {
		w.withArgsLabel = false
	}
}

// WithShortQueryNames changes the query labels on the spans to be shortened versions
// i.e `find_user($1, $2, $3, $4, $5)` --> `find_user`
func WithShortQueryNames(t Tracer) Opt {
	return func(w *wrappedDriver) {
		w.shortQueryName = true
	}
}
