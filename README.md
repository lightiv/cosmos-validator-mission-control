# Validator Mission Control tool

**Validator mission control** tool provides a wide verity of metrics and alerts for validator node operators. It can be installed on validator node directly or any other monitoring nodes (with some special firewall on validator node). We utilized the power of Grafana + Telegraf and extended the monitoring & alerting with custom built go server.

## Install Prerequisites
- **Go 13.x+**
- **Docker 19+**
- **Grafana 6.7+**
- **InfluxDB 1.7+**
- **Telegraf 1.14+**
- **Gaia client**

### A - Install Grafana for Ubuntu
Download the latest .tar.gz file and extract it by using the following commands

```sh
$ cd $HOME
$ wget https://dl.grafana.com/oss/release/grafana-6.7.2.linux-amd64.tar.gz
$ tar -zxvf grafana-6.7.2.linux-amd64.tar.gz
```

Start the grafana server
```sh
$ cd grafana-6.7.2/bin/
$ ./grafana-server

Grafana will be running on port :3000 (ex:: https://localhost:3000)
```

### Install InfluxDB

Download the latest .tar.gz file and extract it by using the following commands

```sh
$ cd $HOME
$ wget https://dl.influxdata.com/influxdb/releases/influxdb-1.7.10_linux_amd64.tar.gz
$ tar xvfz influxdb-1.7.10_linux_amd64.tar.gz
```

Start influxDB

```sh
$ cd $HOME and run the below command to start the server
$ ./influxdb-1.7.10-1/usr/bin/influxd

The default port that runs the InfluxDB HTTP service is :8086
```

**Note :** If you want to give custom configuration then you can edit the `influxdb.conf` at `/influxdb-1.7.10-1/etc/influxdb` and do not forget to restart the server after the changes. You can find a sample 'influxdb.conf' [file here]<https://github.com/jheyman/influxdb/blob/master/influxdb.conf>.


### Install Telegraf

Download the latest .tar.gz file and extract it by using the following commands
```sh
$ cd $HOME
$ wget https://dl.influxdata.com/telegraf/releases/telegraf-1.14.0_linux_amd64.tar.gz
tar xf telegraf-1.14.0_linux_amd64.tar.gz
```

Start telegraph
```sh
$ cd telegraf/usr/bin/
$ ./telegraf --config ../../etc/telegraf/telegraf.conf
```

### Setup a rest-server on validator instance
If your validator instance does not have a rest server running, execute this command to setup the rest server

```sh
gaiacli rest-server --chain-id cosmoshub-3 --laddr tcp://127.0.0.1:1317
```

## Install and configure the Validator Mission Control Tool

### Get the code

```bash
$ git clone git@github.com:chris-remus/chainflow-vitwit.git
$ cd chainflow-vitwit
$ cp example.config.toml config.toml
```

### Configure the following variables in `config.toml`

- *tg_chat_id*

    Telegram chat ID to receive telegram alerts
- *tg_bot_token*

    Telegram bot token. The bot should be added to the chat and it should have send message permission

- *email_address*

    E-mail address to receive mail notifications

- *sendgrid_token*

    Sendgrid mail service api token.
- *missed_blocks_threshold*

    Configure the threshold to receive  **Missed Block Alerting**
- *block_diff_threshold*

    An integer value to receive **Block difference alerts**

- *alert_time1* and *alert_time2*

    These are for regular status updates. To receive validator status daily (twice), configure these parameters in the form of "02:25PM". The time here refers to UTC time.

- *voting_power_threshold*

    Configure the threshold to receive alert when the voting power reaches or drops below of the threshold given.

- *num_peers_threshold*

    Configure the threshold to get an alert if the no.of connected peers falls below the threshold.

- *enable_telegram_alerts*

    Configure **yes** if you wish to get telegram alerts otherwise make it **no** .

- *enable_email_alerts*

    Configure **yes** if you wish to get email alerts otherwise make it **no** .

- *validator_rpc_endpoint*

    Validator rpc end point(RPC of your own validator) useful to gather information about network info, validatr voting power, unconfirmed txns etc.

- *val_operator_addr*

    Operator address of your validator which will be used to get staking, delegation and distribution rewards.

- *account_addr* 

    Your validator account address which will be used to get account informtion etc.

- *validator_hex_addr*

    Validator hex address useful to know about last proposed block, missed blocks and voting power.

- *lcd_endpoint*

    Address of your lcd client (ex: http://localhost:1317)

- *external_rpc*

    External open RPC endpoint(secondary RPC other than your own validator). Useful to gather information like validator caught up, syncing and missed blocks etc

After populating config.toml, build and run the monitoring binary

```bash
$ go build -o chain-monit && ./chain-monit
```

### Run using docker
```bash
$ docker build -t cfv .
$ docker run -d --name chain-monit cfv
```

We have finished the installation and started the server. Lets configure Grafana dashboard.

## Granfana Dashboards

In grafana demo provided, you will see three dashboards

1. Validator Monitoring Metrics (These are the metrics which we have calculated and stored in influxdb)
2. System Metrics (These are related to system configuration which comes from telegraf)
3. Summary (Which gives a quick information about validator and system metrics)


### 1. Validator monitoring metrics
The following list of metrics are displayed in this dashboard.

- Validator Details :  Displays the details of a validator like moniker, website, keybase identity and details.
- Gaiad Status :  Displays whether the gaiad is running or not in the from of UP and DOWN.
- Validator Status :  Displays the validator health. Shows Voting if the validator is in active state or else Jailed.
- Gaiad Version : Displays the version of gaia currently running.
- Validator Caught Up : Displays whether the validator node is in sync with the network or not.
- Block Time Difference : Displays the time difference between previous block and current block.
- Current Block Height :  Validator : Displays the current block height committed by the validator.
- Latest Block Height : Network : Displays the latest block height of a network.
- Height Difference : Displays the difference between heights of validator current block height and network latest block height.
- Last Missed Block Range : Displays the continuous missed blocks range based on the threshold given in the config.toml
- Blocks Missed In last 48h : Displays the count of blocks missed by the validator in last 48 hours.
- Unconfirmed Txns : Displays the number of unconfirmed transactions on that node.
- No.of Peers : Displays the total number of peers connected to the validator.
- Peer Address : Displays the ip addresses of connected peers.
- Latency : Displays the latency of connected peers with respect to the validator.
- Validator Fee : Displays the commission rate of the validator.
- Max Change Rate : Displays the max change rate of the commission.
- Max Rate : Displays the max rate of the commission.
- Voting Power : Displays the voting power of the validator.
- Self Delegation Balance : Displays delegation balance of the validator.
- Current Balance : Displays the account balance of the validator.
- Unclaimed Rewards : Displays the current unclaimed rewards amount of the validator.
- Last proposed Block Height : Displays height of the last block proposed by the validator.
- Last Proposed Block Time : Displays the time of the last block proposed by the validator.
- Voting Period Proposals : Displays the list of the proposals which are currently in voting period.
- Deposit Period Proposals : Displays the list of the proposals which are currently in deposit period.
- Completed Proposals : Displays the list of the proposals which are completed with their status as passed or rejected.


**Note:** Above mentioned metrics will be calculated and displayed according to the validator address which will be configured in config.toml

For alerts regarding system metrics, a telegram bot can be set up on the dashboard itself. A new notification channel can be added for telegram bot by clicking on the bell icon on the left hand sidebar of the dashboard. This will let the user configure the telegram bot id and chat id. A custom alert can be set for each graph by clicking on the edit button and adding alert rules.

### 2. System Monitoring Metrics
These are powered by telegraf and we are utilizing the features as is. We don't see a need to extend these metrics as these are highly informative and granular.
-  For the list of system monitoring metrics, you can refer `telgraf.conf`. You can just replace it with your original telegraf.conf file which will be located at /telegraf/etc/telegraf (installation directory)
 
 ### 3. Summary Dashboard
This dashboard displays a quick information summary of validator details and system metrics. It has following details.

- Validator identity (Moniker, Website, Keybase Identity, Details, Operator Address and Account Address), validator summary (Gaiad Status, Validator Status, Voting Power, Height Difference and No.Of peers) are the metrics being displayed from Validator details.

- CPU usage, RAM Usage, Memory usage and information about disk usage are the metrics being displayed from System details.

## How to import these dashboards in your grafana?

### 1. Login to your grafana dashboard
- Open your web browser and go to http://<your_ip>:3000/. `3000` is the default HTTP port that Grafana listens to if you haven’t configured a different port.
- If you are a first time user type `admin` for the username and password in the login page.
- You can change the password after login.

### Import the dashboards
- To import the json file of the **validator monitoring metrics** click the *plus* button present on left hand side of the dashboard. Click on import and load the validator_monitoring_metrics.json present in the grafana_template folder. 

- Select the datasources and click on import.

- To import **system monitoring metrics** click the *plus* button present on left hand side of the dashboard. Click on import and load the system_monitoring_metrics.json present in the grafana_template folder

- While creating this dashboard if you face any issues at valueset, change it to empty and then click on import by selecting the datasources.

- To import **summary**, you can just follow the above steps which you did for validator monitoring metrics or system monitoring metrics. To import the json you can find the summary.json template in grafana_template folder.

- *For more info about grafana dashboard imports you can refer https://grafana.com/docs/grafana/latest/reference/export_import/*


## Alerting (Telegram and Email)
 We have developed a custom alerting module to alert on several events. It uses the data from influxdb and trigger the alerts based on user configured thresholds.

 - Alert about validator node sync status.
 - Alert when missed blocks when the missed blocks count reaches or exceedes **missed_blocks_threshold** which is user configured in *config.toml*
 - Alert when no.of peers when the count falls below of **num_peers_threshold** which is user configured in *config.toml*
- Alert when the block difference between network and validator reaches or exceeds the **block_diff_threshold** which is user configured in *config.toml*
- Alert when the gaiad status is not running on validator instance.
- Alert when a new proposal is created.
- Alert when the proposal is moved to voting period, passed or rejected.
- Alert when voting period proposals is due in less than 24 hours and also if the validator didn't vote on proposal yet.
- Alert about validator health whether it's voting or jailed. You can get alerts twice a day based on the time you have configured **alert_time1** and **alert_time2** in *config.toml*
- Alert when the voting power of your validator drops below **voting_power_threshold** which is user configured in *config.toml*


## Hosting it on separate monitoring node

This monitoring tool can also be hosted on any separate monitoring node/public sentry node of the validator.

 - Prerequisites and setup for sentry node remains the same with 1 exception. Telegraf should be installed on the validator instance instead of sentry node. Telegraf collects all the hardware matrics like CPU load, RAM usage etc and post it to InfluxDB. InfluxDB sits on monitoring node. 
 - While importing and setting up the dashboards for Grafana, the url has to be changed for InfluxDBTelegraf datasource to point to validator node telegraf url.
 - As mentioned above, the default port on which Telegraf points the data is 8086, so the url should be replaced as "http://validator-ip:8086"
 - This will allow Grafana to display system metrics of validator instance instead of displaying metrics of monitoring node.
 - All other metrics can be colelcted from monitoring node itself. Monitoring node should have the access to validator RPC and LCD endpoints. Configure your validator firewall to allow Monitoring node to access these endpoints.
 

**Note:**
We thought well about the validator operator needs and covered all the required features for monitoring and alerting. Please feel free to create issues if you think of any missing feature.