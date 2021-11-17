package runner

import "github.com/projectdiscovery/gologger"

const banner = `
       __             _____              
      / /_  __  _____/__  /_______  _  __
     / __ \/ / / / __ \/ / ___/ _ \| |/_/
    / / / / /_/ / / / / / /  /  __/>  <  
   /_/ /_/\__,_/_/ /_/_/_/   \___/_/|_|	  v0.0.1
`

// Version is the current version of hun7rex
const Version = `0.0.1`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\t0x71rex\n\n")

	gologger.Print().Msgf("Use with caution. You are responsible for your actions\n")
	gologger.Print().Msgf("Developers assume no liability and are not responsible for any misuse or damage.\n")
}
