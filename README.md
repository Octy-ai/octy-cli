# octy-cli
<img  src="https://octy.ai/images/logo/text_side_trans@2x.png"  alt="drawing"  width="100"/>
A command-line tool for Octy
  

## Get started with the Octy CLI

Manage your account configurations, algorithm configurations, raw data resources and Octy object definition resources from your terminal.

The Octy CLI is a developer tool to help you manage your integration with Octy directly from your terminal. It’s simple to install, works on macOS, Windows, and Linux. 
With the Octy CLI you can:

- Update your account configurations
- Update Octy algorithm configurations
-   Create any of the following raw data resources: *profiles, items, event instances*
-   Create, retrieve, update*, or delete any of the following Octy object definition resources: *event types, segments, message templates* *updates can only be applied to message templates
-   Generate a churn prediction analysis report

## Installation

Octy CLI is available for macOS, Windows, and Linux for distros like Ubuntu, Debian, RedHat and CentOS.

You can download the binary executable without having to install and additional dependencies. Find our  [latest release](https://github.com/Octy-ai/octy-cli/releases/latest)  and download the  `octy-cli_X.X.X_mac-os`,  `octy-cli_X.X.X_windows`, or  `octy-cli_X.X.X_linux`  `tar.gz`  file, unzip it, and inside you'll have the  `octy-cli`  executable that you can run directly.

On macOS or Linux, you can move this file to  `/usr/local/bin`  or  `/usr/bin`  locations to have it be runnable from anywhere. Otherwise,  `cd`  to the folder where you unzipped the  `tar.gz`  file and invoke the CLI with  `./octy-cli`.

## Updating

When a new version is available, the Octy CLI will notify you with each command that is run, providing you with a link to your operating systems relevant latest release. On macOS or Linux, you can move this file to  `/usr/local/bin`  or  `/usr/bin`  locations to replace the current CLI and have it be runnable from anywhere. 

## Usage

Installing the CLI provides access to the `octy-cli` command.
```sh-session
./octy-cli [command]

# Run `--help` for detailed information about CLI commands
./octy-cli [command] help
```

## Commands

#### Top level commands:
- `auth`
- `apply`
- `upload`
- `get`
- `delete`

#### Command references:

Authentication
- [`auth --public-key pk_xxx --secret-key sk_xxx` ](#auth)

Configurations
- [`apply --filepath path/to/account-configurations.yml` ](#update-account-configurations)
- [`get accountconfig` ](#get-account-configurations)
- [`apply --filepath path/to/algorithm-configurations.yml` ](#update-algorithm-configurations)
- [`get algorithmconfig <type>` ](#get-algorithm-configurations)

Event types
- [`apply --filepath path/to/event-type-definitions.yml` ](#create-event-types)
- [`get eventtypes` ](#get-event-types)
- [`delete eventtypes <event type id>` ](#delete-event-types)

Segments
- [`apply --filepath path/to/segment-definitions.yml` ](#create-segments)
- [`get segments` ](#get-segments)
- [`delete segments <segment id>` ](#delete-segments)

Templates
- [`apply --filepath path/to/message-template-definitions.yml` ](#create-message-templates)
- [`get templates` ](#get-message-templates)
- [`delete templates <template id>` ](#delete-message-templates)

Profiles
- [`upload profiles --filepath path/to/profiles.csv` ](#create-profiles)

Items
- [`upload items --filepath path/to/items.csv` ](#create-items)

Event instances
- [`upload events --filepath path/to/events.csv` ](#create-events)

Churn prediction report
- [`get churnreport` ](#get-churn-prediction-report)


# auth
Configure your Octy API Authentication credentials, granting this CLI limited access to your Octy account within the scope of the available options.
Your Octy credentials will be stored in your devices OS keychain.

**flags**
`--public-key`  `-p` 			(Your Octy public key **required**)
`--secret-key`  `-s` 			(Your Octy secret key **required**)

**example cmd**
`./octy-cli auth -p pk_xxx -s sk_xxx`


# update account configurations
Update your Octy account configurations. Go [here](https://octy.ai/docs/api#AccSetConfigurations) for more info on account configurations.

**flags**
`--filepath`  `-f`			(Path to account configurations yaml file **required**)

**example cmd**
`./octy-cli apply -f path/to/account-configurations.yml`

**example account-configurations.yml**
```
kind: accountConfigurations
configurations :
	contactName: Ben
	contactSurname: Goodenough
	contactEmail: support@octy.ai
	webhookURL: https://octy.ai/hook/1
	authenticatedIDKey : shop-id-key
```


# get account configurations
Retrieve your current Octy account configurations.

**example cmd**
`./octy-cli get accountconfig`


# update algorithm configurations
Update your Octy algorithm configurations. Go [here](https://octy.ai/docs/api#AlgoSetConfigurations) for more info on algorithm configurations.

**flags**
`--filepath`  `-f`			(Path to algorithm configurations yaml file **required**)

**example cmd**
`./octy-cli apply -f path/to/algorithm-configurations.yml`

**example algorithm-configurations.yml**
```
kind: algorithmConfigurations
configurations :

	- algorithmName: churn
	  configurations:
		profileFeatures:
			- visits
			- balance

	- algorithmName: rec
	  configurations:
		recommendInteractedItems: true
		itemIDStopList:
			- product-1234
		profileFeatures:
			- visits
			- balance
```


# get algorithm configurations
Retrieve your current Octy algorithm configurations.

**example cmd**
Return all algorithm configurations
`./octy-cli get algorithmconfig`
Return recommendations algorithm configurations only
`./octy-cli get algorithmconfig rec`
Return churn prediction analysis algorithm configurations only
`./octy-cli get algorithmconfig churn`


# create event types
Create custom event type definitions. Go [here](https://octy.ai/docs/creating_resources#Custom%20and%20System%20event%20types) for more on custom event type definitions.

**flags**
`--filepath`  `-f`			(Path to event type definitions yaml file **required**)

**example cmd**
`./octy-cli apply -f path/to/event-type-definitions.yml`

**example event-type-definitions.yml**
```
kind: eventTypes
eventTypeDefinitions :

- eventType: login
  eventProperties:
	- device
	- time
- eventType: logout
  eventProperties:
	- device
```

# get event types
Retrieve event type definitions. You can specify up to 100 event type identifiers with each command.

**flags**
`--ids`  `-i`			(Only output the identifiers of returned event type definitions from the API [true/false] **optional**)

**example cmd**
Return all custom event type definitions
`./octy-cli get eventtypes`
Return custom event type definitions where event_type_id equals `custom_event_type_5cc1b718` or `custom_event_type_036ea654`
`./octy-cli get eventtypes custom_event_type_5cc1b718 custom_event_type_036ea654`
Return all custom event type identifiers
`./octy-cli get eventtypes --ids=true`
Return custom event type identifiers where event_type_id equals `custom_event_type_5cc1b718` or `custom_event_type_036ea654`
`./octy-cli get eventtypes custom_event_type_5cc1b718 custom_event_type_036ea654 --ids=true`


# delete event types
Delete specified custom event type definitions.  You can specify up to 100 event type identifiers with each command.

**example cmd**
`./octy-cli delete eventtypes custom_event_type_5cc1b718`


# create segments
Create segment definitions. Go [here](https://octy.ai/docs/segmentation#Create%20segment%20definitions) for more on creating segment definitions.

**flags**
`--filepath`  `-f`			(Path to segment definitions yaml file **required**)

**example cmd**
`./octy-cli apply -f path/to/segment-definitions.yml`

**example segment-definitions.yml**
```
kind: segments
segmentDefinitions:
# Customers who login and make a purchase within 5 minutes
- segmentName: fast paying customers
  segmentType: live
  segmentSubtype: 2 
  segmentTimeframe: 0
  eventSequence:
	- eventType: login
	  expTimeframe: 5
	  actionInaction: action
	- eventType: charged
	  expTimeframe: 0
	  actionInaction: inaction
	  eventProperties:
		item_id : ios
		payment_method : visa
  profilePropertyName:
  profilePropertyValue:
  
# Customers that have accessed your systems but made no purchase in the past 60 days
- segmentName: worst customers  
  segmentType: past
  segmentSubtype: 2
  segmentTimeframe: 60
  eventSequence:
  - eventType: login
    expTimeframe: 0
    actionInaction: action
  - eventType: charged
    expTimeframe: 0
    actionInaction: inaction
profilePropertyName: customer_tier
profilePropertyValue: gold
```

# get segments
Retrieve segment definitions. You can specify up to 100 segment identifiers with each command.

**flags**
`--ids`  `-i`			(Only output the identifiers of returned segment definitions from the API [true/false] **optional**)

**example cmd**
Return all segment definitions
`./octy-cli get segments`
Return segment definitions where segment_id equals `segment_5cc1b718` or `segment_036ea654`
`./octy-cli get segments segment_5cc1b718 segment_036ea654`
Return all segment identifiers
`./octy-cli get segments --ids=true`
Return segment identifiers where segment_id equals `segment_5cc1b718` or `segment_036ea654`
`./octy-cli get segments segment_5cc1b718 segment_036ea654 --ids=true`


# delete segments
Delete specified segment definitions.  You can specify up to 100 segment identifiers with each command.

**example cmd**
`./octy-cli delete segments segment_036ea654`


# create message templates
Create message template definitions. Go [here](https://octy.ai/docs/messaging#Create%20message%20templates) for more on creating message template definitions.

**flags**
`--filepath`  `-f`			(Path to message template definitions yaml file **required**)

**example cmd**
`./octy-cli apply -f path/to/message-template-definitions.yml`

**example message-template-definitions.yml**
```
kind: templates
templateDefinitions :

- friendlyName: Basic personalised greeting
  templateType: email
  title : We have some great deals for you!
  content : Hi {first_name}, we have some great deals for you {greeting}!
  defaultValues:
	first_name: there
	greeting : my friend
  templateID : template_1234  # As a template ID is supplied, cli will attempt to update this template

- friendlyName: Website heading banner
  templateType: website
  title : --
  content : Wait! before you think about going...
  defaultValues:
  metadata: 
	# show this banner to customers with a high churn probability and a high RFM score, i.e your best customers that are likely to churn
	churn_pred : high
	rfm_score_upper : 444
	rfm_score_lower : 333
```

# get message templates
Retrieve message templates definitions. You can specify up to 100 template identifiers with each command.

**flags**
`--ids`  `-i`			(Only output the identifiers of returned message template definitions from the API [true/false] **optional**)

**example cmd**
Return all message template definitions
`./octy-cli get templates`
Return message template definitions where template_id equals `template_5cc1b718` or `template_036ea654`
`./octy-cli get templates template_5cc1b718 template_036ea654`
Return all message template identifiers
`./octy-cli get templates --ids=true`
Return message template identifiers where segment_id equals `template_5cc1b718` or `template_036ea654`
`./octy-cli get templates template_5cc1b718 template_036ea654 --ids=true`


# delete message templates
Delete specified message template definitions.  You can specify up to 100 template identifiers with each command.

**example cmd**
`./octy-cli delete templates template_036ea654`


# create profiles
Create profiles. Go [here](https://octy.ai/docs/creating_resources#Creating%20profiles) for more on creating profiles.

**flags**
`--filepath`  `-f`			(Path to profiles csv (comma separated value) file **required**)

**example cmd**
`./octy-cli upload profiles -f path/to/profiles.csv`

**example profiles.csv**
Download an example profiles.csv file [here](https://raw.githubusercontent.com/Octy-ai/octy-cli/master/examples/upload_cmd_csvs/profiles.csv)
|  customer_id | has_charged | profile_data>>likes | profile_data>>gender | profile_data>>age | platform_info>>os | 
|--|--|--|--|--|--|
| 2748275927498 | true | 4300 | female | 28 | ios |
| 4438286378150 | true | 5600 | female | 24 | android |
| 2714793653064 | false | 2467 | other | 33 | macOS |


# create items
Create items. Go [here](https://octy.ai/docs/creating_resources#Creating%20items) for more on creating items.

**flags**
`--filepath`  `-f`			(Path to items csv (comma separated value) file **required**)

**example cmd**
`./octy-cli upload items -f path/to/items.csv`

**example items.csv**
Download an example items.csv file [here](https://raw.githubusercontent.com/Octy-ai/octy-cli/master/examples/upload_cmd_csvs/items.csv)
|  item_id | item_category | item_name | item_description | item_price | 
|--|--|--|--|--|
| 01134f95-d1a8 | Clothing | T-Shirt | Super cool black t-shirt | 2000 |
| c1600e7b-84cf | Clothing | Jeans | Awesome blue jeans | 4900 |
| dc58879e-5fca | Outdoor | Tent | All-weather outdoor tent | 6290 |


# create events
Create events. Go [here](https://octy.ai/docs/creating_resources#Creating%20events) for more on creating events.

**flags**
`--filepath`  `-f`			(Path to events csv (comma separated value) file **required**)

**example cmd**
`./octy-cli upload events -f path/to/events.csv`

**example events.csv**
Download an example events.csv file [here](https://raw.githubusercontent.com/Octy-ai/octy-cli/master/examples/upload_cmd_csvs/events.csv)
|  event_type | profile_id | created_at | event_properties>>item_id | event_properties>>payment_method | 
|--|--|--|--|--|
| charged | profile_1234 | 2021-06-29 18:26:44 | 01134f95-d1a8 | applepay |
| charged | profile_1234 | 2021-08-24 23:25:49 | c1600e7b-84cf | visa |
| charged | profile_1234 | 2021-06-11 11:57:57 | dc58879e-5fca | visa |


# get churn prediction report
Generate and retrieve current churn prediction report, detailing the key features that lead to customer churn.
Go [here](https://octy.ai/docs/api#ChurnReportObject) for more on the churn prediction report.

**flags**
`--outpath`  `-o`			(Path to a directory where a markdown file containing a churn report will be saved **optional**)

**example cmd**
`./octy-cli get churnreport -o path/to/save/churn-report/`



# License

Copyright (c) Octy LTD. All rights reserved.
Licensed under the <a  href="https://www.apache.org/licenses/LICENSE-2.0.txt">Apache License 2.0 license</a>.

 
# Author

Ben Goodenough,
CEO & Founder Octy LTD
