package lib

// MetricRequestMatcher ...
type MetricRequestMatcher struct {
	Metrics []Metric
}

// NewMetricRequestMatcherFromConfig ...
func NewMetricRequestMatcherFromConfig(config *ActionsWorkerConfig) *MetricRequestMatcher {
	metrics := []Metric{}

	for _, metric := range config.MetricsGroup.Metrics {
		dimensions := []Dimension{}

		for _, dimension := range metric.DimensionsGroup.Dimensions {
			dimensions = append(dimensions, dimension)
		}

		metrics = append(metrics, Metric{Dimensions: dimensions})
	}

	return &MetricRequestMatcher{
		Metrics: metrics,
	}
}

// Matching ...
func (mg *MetricRequestMatcher) Matching(trackRequest *TrackRequest) []Metric {
	trackRequestMap := trackRequest.CreateMap()
	matchingMetrics := []Metric{}

metrics:
	for _, metric := range mg.Metrics {
		for _, dimension := range metric.Dimensions {
			if _, ok := trackRequestMap[dimension.Name]; !ok {
				continue metrics
			}

			if !dimension.Condition.Meets(trackRequestMap[dimension.Name]) {
				continue metrics
			}
		}

		matchingMetrics = append(matchingMetrics, metric)
	}

	return matchingMetrics
}
