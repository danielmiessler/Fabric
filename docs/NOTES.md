## Notes on some refactoring.

- The goal is to bring more encapsulation of the models management and simplified configuration management to bring increased flexibility, transparency on the overall flow, and simplicity in adding new model.
- We need to differentiate:
  - Vendors: the producer of models (like OpenAI, Azure, Anthropic, Ollama, ..etc) and their associated APIs
  - Models: the LLM models these vendors are making public
- Each vendor and operations allowed by the vendor needs to be encapsulated. This includes:
  - The questions needed to setup the model (like the API key, or the URL)
  - The listing of all models supported by the vendor
  - The actions performed with a given model

- The configuration flow works like this for an **initial** call:
  - The available vendors are called one by one, each of them being responsible for the data they collect. They return a set of environment variables under the form of a list of strings, or an empty list if the user does not want to setup this vendor. As we do not want each vendor to know which way the data they need will be collected (e.g., read from the command line, or a GUI), they will be asked for a list of questions, the configuration will inquire the user, and send back the questions with the collected answers to the Vendor. The Vendor is then either instantiating an instance (Vendor configured) and returning it, or returning `nil` if the Vendor should not be set up.
  - the `.env` file is created, using the information returned by the vendors
  - A list of patterns is downloaded from the main site

- When the system is configured, the configuration flows:
  - Read the `.env` file using the godotenv library
  - It configures a structure that contains the various vendors selected as well as the preferred model. This structure will be completed with some of the command line values (i.e, context, session, etc..)

- To get the list of all supported models:
  - Each configured model (part of the configuration structure) is asked, using a goroutine, to return the list of model

- Order when building message: session + context + pattern + user input (role "user)


## TODO:
- Check if we need to read the system.md for every patterns when running the ListAllPatterns
- Context management seems more complex than the one in the original fabric. Probably needs some work (at least to make it clear how it works)
- models on command line: give as well vendor (like `--model openai/gpt-4o`). If the vendor is not given, get it by retrieving all possible models and searching from that.
- if user gives the ollama url on command line, we need to update/init an ollama vendor.
- The db should host only things related to access and storage in ~/.config/fabric
- The interaction part of the Setup function should be in the cli (and perhaps all the Setup)
