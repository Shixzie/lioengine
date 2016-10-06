package lioengine

import (
	"errors"
	"log"
	"net/http"
	"sync"
)

// Bot holds all utility for searching news and updates
type Bot struct {
	currentProviders []*provider
}

// NewBot creates a bot.
func NewBot() (bot *Bot) {
	bot = new(Bot)
	bot.currentProviders = []*provider{}
	return
}

// FindUpdatesFor is the 'main' function for this package. And will return
// an []*Update for the given project name.
func (b *Bot) FindUpdatesFor(projectName string) (updates []*Update, err error) {
	err = b.makeAPICalls(projectName)
	if err != nil {
		log.Printf("Error ocurred at lioengine.go - makeApiCalls(...) : %s", err.Error())
		return
	}
	//updates, err = analizeUpdates()
	// if err != nil {
	// 	log.Printf("Error ocurred at lioengine.go - analizeUpdates(...) : %s", err.Error())
	// 	return
	// }
	return
}

// makeApiCalls connects to the apiServer, and fetch all the results
// for later ai ml algorithms.
func (b *Bot) makeAPICalls(projectName string) (err error) {
	var wg = new(sync.WaitGroup)
	for _, provider := range b.currentProviders {
		if provider.RequestInfo.Request == nil {
			log.Println("Nil request on makeAPICall")
			provider.RequestInfo.Request, err = http.NewRequest("", parseURL(provider.RequestInfo.urlWithParameters, projectName, provider.RequestInfo.Quantity), nil)
			if err != nil {
				log.Printf("Error ocurred at requests.go - http.NewRequest(...) : %s", err.Error())
				return
			}
		}

		switch v := provider.Type.(type) {
		case bing:
			wg.Add(1)
			go v.search(provider, wg)
			break
		}
	}
	wg.Wait()
	return
}

// AddUpdatesProvider adds the news provider by the name given and
// initializes it with the corresponding apiToken.
// This function should be called before FindUpdatesFor().
// This is also designed to be called multiple times.
// Current supported providers are: Bing.
func AddUpdatesProvider(newProviderName, apiToken string, bots ...*Bot) (err error) {
    //Iterates through all bots to add the providers
	for _, bot := range bots {
		var alreadyAdded = false
		// Iterates through all currentProviders
		for _, currentProvider := range bot.currentProviders {
			// Checks if we already have this provider
			if newProviderName == currentProvider.Name {
				alreadyAdded = true
			}
		}

		// If we do, then return
		if alreadyAdded {
			err = errors.New("This provider is already added.")
			return
		}

		var itsASupportedProvider = false
		// Iterates through all our supported providers
		for _, supportedProviderName := range supportedProviders {
			if newProviderName == supportedProviderName {
				itsASupportedProvider = true
			}
		}

		// If is one of our supported providers and we haven't added it yet, then we add it.
		if itsASupportedProvider {
			bot.setupProvider(newProviderName, apiToken)
		} else { // If provider not supported.
			err = errors.New("This provider is not supported by the bot.")
			return
		}
	}
	return
}

// setupProvider generates a provider corresponding to it's name
func (b *Bot) setupProvider(providerName, apiToken string) {
	switch providerName {
	case "Bing":
		bing := bing{}
		provider := bing.setup(apiToken)
		b.currentProviders = append(b.currentProviders, &provider)
		return
	}
	return
}

// Returns the slide index for the provider with the name
// providerName.
func (b *Bot) getProviderIndexByName(providerName string) (index int) {
	for index, currentProvider := range b.currentProviders {
		if currentProvider.Name == providerName {
			return index
		}
	}
	return
}