package emit

type option func(*config)

func withBuffSize(size int) option {
  return func(c *config) {
    if size >= 1 {
      c.size = size
    }
  }
}

func withStrategy(s strategy) option {
  return func(c *config) { c.strategy = s }
}
