
## dolores

> name reference: Westworld series
> dir structure is similar to the concept of consciousness in there
>
> * corecode : overriding `drive` for all hosts; here core helper code
>
> * loops : the general narrative that bot stick to; here for presence on Slack, etc
>
> * drives : the goal to host's actions; here for Actions driven/demanded by presence in `loops`
>
> * memories : despite wipes remembering past; here bookkeeping for audits
>
> * reveries : lifelike gestures of hosts; here for scope of `ML` into healing/alerting factors

This bot yet need to find the inner maze.

---

### How conversation flow happens currently

* It will get started and join comm channels (Slack as of now)

* User can say `help` to Dolores for set of command-syntax it responds to. Yeah it's not to play friendly catch-up, yet.

* When any command (even 'help') is recieved by Dolores, they get passed to `loops` section (currently directly to event handler of loops/slack). It checks if message is meant for it and then proceeds.

* In `loops` section, recieved message is processed and identified for which message-handler it is meant.

* The message-handler extracts variable query from message and passes it to relevant `drive` handler for it.

* These `drives` prepare the provided query variables to usable form and call the baseline `corecode` which can help to fetch relevant information.

* The result of all this when returns back to `loops` message-handler, gets sent back as reply to user.

---

