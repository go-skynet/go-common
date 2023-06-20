package common

// Note: These intentionally are taken from go-llama.cpp initially!
// This should generally be a _superset_ of all allowed configuration options.
// A backend can safely _ignore_ something!
// As I march though the go-* projects, I will add anything that is currently missing

type InitializationOptions struct {
	ContextSize int
	Seed        int
	NBatch      int
	F16Memory   bool
	MLock       bool
	MMap        bool
	Embeddings  bool
	NGPULayers  int
	MainGPU     string
	TensorSplit string
}

type PredictTextOptions struct {
	Seed, Threads, Tokens, TopK, Repeat, Batch, NKeep int
	TopP, Temperature, Penalty                        float64
	F16KV                                             bool
	DebugMode                                         bool
	StopPrompts                                       []string
	IgnoreEOS                                         bool

	TailFreeSamplingZ float64
	TypicalP          float64
	FrequencyPenalty  float64
	PresencePenalty   float64
	Mirostat          int
	MirostatETA       float64
	MirostatTAU       float64
	PenalizeNL        bool
	LogitBias         string
	TokenCallback     func(string) bool

	PathPromptCache             string
	MLock, MMap, PromptCacheAll bool
	PromptCacheRO               bool
	MainGPU                     string
	TensorSplit                 string
}

// ==== Setters ====

type PredictTextOptionSetter func(p *PredictTextOptions)

type InitializationOptionSetter func(p *InitializationOptions)

// SetContext sets the context size.
func SetContext(c int) InitializationOptionSetter {
	return func(p *InitializationOptions) {
		p.ContextSize = c
	}
}

func SetModelSeed(c int) InitializationOptionSetter {
	return func(p *InitializationOptions) {
		p.Seed = c
	}
}

// SetContext sets the context size.
func SetMMap(b bool) InitializationOptionSetter {
	return func(p *InitializationOptions) {
		p.MMap = b
	}
}

// SetNBatch sets the  n_Batch
func SetNBatch(n_batch int) InitializationOptionSetter {
	return func(p *InitializationOptions) {
		p.NBatch = n_batch
	}
}

// Set sets the tensor split for the GPU
func SetTensorSplit(maingpu string) InitializationOptionSetter {
	return func(p *InitializationOptions) {
		p.TensorSplit = maingpu
	}
}

// SetMainGPU sets the main_gpu
func SetMainGPU(maingpu string) InitializationOptionSetter {
	return func(p *InitializationOptions) {
		p.MainGPU = maingpu
	}
}

// SetPredictionTensorSplit sets the tensor split for the GPU
func SetPredictionTensorSplit(maingpu string) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.TensorSplit = maingpu
	}
}

// SetPredictionMainGPU sets the main_gpu
func SetPredictionMainGPU(maingpu string) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.MainGPU = maingpu
	}
}

var EnableEmbeddings InitializationOptionSetter = func(p *InitializationOptions) {
	p.Embeddings = true
}

var EnableF16Memory InitializationOptionSetter = func(p *InitializationOptions) {
	p.F16Memory = true
}

var EnableF16KV PredictTextOptionSetter = func(p *PredictTextOptions) {
	p.F16KV = true
}

var Debug PredictTextOptionSetter = func(p *PredictTextOptions) {
	p.DebugMode = true
}

var EnablePromptCacheAll PredictTextOptionSetter = func(p *PredictTextOptions) {
	p.PromptCacheAll = true
}

var EnablePromptCacheRO PredictTextOptionSetter = func(p *PredictTextOptions) {
	p.PromptCacheRO = true
}

var EnableMLock InitializationOptionSetter = func(p *InitializationOptions) {
	p.MLock = true
}

var IgnoreEOS PredictTextOptionSetter = func(p *PredictTextOptions) {
	p.IgnoreEOS = true
}

// SetMlock sets the memory lock.
func SetMlock(b bool) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.MLock = b
	}
}

// SetMemoryMap sets memory mapping.
func SetMemoryMap(b bool) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.MMap = b
	}
}

// SetGPULayers sets the number of GPU layers to use to offload computation
func SetGPULayers(n int) InitializationOptionSetter {
	return func(p *InitializationOptions) {
		p.NGPULayers = n
	}
}

// SetTokenCallback sets the prompts that will stop predictions.
func SetTokenCallback(fn func(string) bool) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.TokenCallback = fn
	}
}

// SetStopWords sets the prompts that will stop predictions.
func SetStopWords(stop ...string) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.StopPrompts = stop
	}
}

// SetSeed sets the random seed for sampling text generation.
func SetSeed(seed int) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.Seed = seed
	}
}

// SetThreads sets the number of threads to use for text generation.
func SetThreads(threads int) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.Threads = threads
	}
}

// SetTokens sets the number of tokens to generate.
func SetTokens(tokens int) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.Tokens = tokens
	}
}

// SetTopK sets the value for top-K sampling.
func SetTopK(topk int) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.TopK = topk
	}
}

// SetTopP sets the value for nucleus sampling.
func SetTopP(topp float64) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.TopP = topp
	}
}

// SetTemperature sets the temperature value for text generation.
func SetTemperature(temp float64) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.Temperature = temp
	}
}

// SetPathPromptCache sets the session file to store the prompt cache.
func SetPathPromptCache(f string) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.PathPromptCache = f
	}
}

// SetPenalty sets the repetition penalty for text generation.
func SetPenalty(penalty float64) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.Penalty = penalty
	}
}

// SetRepeat sets the number of times to repeat text generation.
func SetRepeat(repeat int) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.Repeat = repeat
	}
}

// SetBatch sets the batch size.
func SetBatch(size int) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.Batch = size
	}
}

// SetKeep sets the number of tokens from initial prompt to keep.
func SetNKeep(n int) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.NKeep = n
	}
}

// SetTailFreeSamplingZ sets the tail free sampling, parameter z.
func SetTailFreeSamplingZ(tfz float64) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.TailFreeSamplingZ = tfz
	}
}

// SetTypicalP sets the typicality parameter, p_typical.
func SetTypicalP(tp float64) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.TypicalP = tp
	}
}

// SetFrequencyPenalty sets the frequency penalty parameter, freq_penalty.
func SetFrequencyPenalty(fp float64) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.FrequencyPenalty = fp
	}
}

// SetPresencePenalty sets the presence penalty parameter, presence_penalty.
func SetPresencePenalty(pp float64) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.PresencePenalty = pp
	}
}

// SetMirostat sets the mirostat parameter.
func SetMirostat(m int) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.Mirostat = m
	}
}

// SetMirostatETA sets the mirostat ETA parameter.
func SetMirostatETA(me float64) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.MirostatETA = me
	}
}

// SetMirostatTAU sets the mirostat TAU parameter.
func SetMirostatTAU(mt float64) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.MirostatTAU = mt
	}
}

// SetPenalizeNL sets whether to penalize newlines or not.
func SetPenalizeNL(pnl bool) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.PenalizeNL = pnl
	}
}

// SetLogitBias sets the logit bias parameter.
func SetLogitBias(lb string) PredictTextOptionSetter {
	return func(p *PredictTextOptions) {
		p.LogitBias = lb
	}
}

// ==== Mergers ====

// Create a new PredictTextOptions object with the given options.
func MergePredictTextOptions(p *PredictTextOptions, opts ...PredictTextOptionSetter) *PredictTextOptions {
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func GetMergePredictTextOptionsFnFromDefault(defaultOptions PredictTextOptions) func(opts ...PredictTextOptionSetter) *PredictTextOptions {
	return func(opts ...PredictTextOptionSetter) *PredictTextOptions {
		optionCopy := defaultOptions
		return MergePredictTextOptions(&optionCopy, opts...)
	}
}

// Create a new InitializationOptions object with the given options.
func MergeInitializationOptions(p *InitializationOptions, opts ...InitializationOptionSetter) *InitializationOptions {
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Explicitly pass by value to copy the default rather than modify it.
func GetMergeInitializationOptionsFnFromDefault(defaultOptions InitializationOptions) func(opts ...InitializationOptionSetter) *InitializationOptions {
	return func(opts ...InitializationOptionSetter) *InitializationOptions {
		optionCopy := defaultOptions
		return MergeInitializationOptions(&optionCopy, opts...)
	}
}
