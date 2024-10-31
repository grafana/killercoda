# Gameplay Instructions

- Upon starting the game, you will receive a description of your current location and a list of available actions.

- Type the command corresponding to the action you want to take and press **Enter**.

- The game continues based on your inputs and choices.

- This game involves checking Grafana dashboards to progress. You can access the Grafana dashboard at [http://localhost:3000]({{TRAFFIC_HOST1_3000}}). Check the dashboard for hints and clues.

## Available Commands

At any point, you can type `list actions`{{copy}} to see the available commands in your current location.

Some universal commands include:

- `quit`{{copy}} or `exit`{{copy}}: End the game.

- `list actions`{{copy}}: Display available actions.

**Sample Actions**:

- **Movement**:
  - `go north`{{copy}}

  - `go south`{{copy}}

  - `go to town`{{copy}}

- **Interactions**:
  - `request sword`{{copy}}

  - `pick herb`{{copy}}

  - `explore`{{copy}}

  - `accept quest`{{copy}}

  - `look at sword`{{copy}}

  - `pray`{{copy}}

- **Special Commands**:
  - `cheat`{{copy}} (to obtain a sword immediately; not recommended).

_Note_: Not all actions are available in every location. Some actions may require certain conditions to be met or prerequisites to be fulfilled.

## Tips for Playing

- **Explore Thoroughly**: Don’t hesitate to try different actions to discover hidden elements.

- **Manage Your Items**: Keep track of items like swords and herbs; they can affect your interactions.

- **Interact with Characters**: Talking to NPCs like the blacksmith, wizard, or priest can open new paths.

- **Monitor Forge Heat**: When at the blacksmith, you’ll need to manage the forge’s heat to get your sword.

- **Beware of Choices**: Some decisions, like accepting the wizard’s offer, have consequences.

## Sample Gameplay Flow

1. **Starting Out**:

   - You’re at the starting point with the option to `go north`{{copy}} or `cheat`{{copy}}.

   - Typing `go north`{{copy}} takes you to the forest.

1. **In the Forest**:

   - Options include `go north`{{copy}} to the cave, `go south`{{copy}} back to start, `go to town`{{copy}}, or `pick herb`{{copy}}.

   - You might choose to `pick herb`{{copy}} and then `go to town`{{copy}}.

1. **In the Town**:

   - Several locations to explore: `blacksmith`{{copy}}, `mysterious man`{{copy}}, `quest giver`{{copy}}, `chapel`{{copy}}.

   - Visit the `blacksmith`{{copy}} to `request sword`{{copy}}.

1. **At the Blacksmith**:

   - After requesting a sword, you’ll need to `heat forge`{{copy}} and `check sword`{{copy}} periodically.

   - Adjust the forge heat using `heat forge`{{copy}} and `cool forge`{{copy}} until the sword is ready.

1. **Getting the Sword**:

   - Once the forge is at the correct temperature, `check sword`{{copy}} will let you obtain it.

   - With the sword, you can interact with other characters differently.

1. **Meeting the Wizard**:

   - Return to town and choose `mysterious man`{{copy}} to meet the wizard (requires having a sword).

   - Decide whether to `accept his offer`{{copy}} or `decline his offer`{{copy}}.

1. **Accepting a Quest**:

   - Visit the `quest giver`{{copy}} to `accept quest`{{copy}}.

   - Your ability to complete the quest may depend on previous choices.

1. **Visiting the Chapel**:

   - Go to the `chapel`{{copy}} and `look at sword`{{copy}} to interact with the priest.

   - The priest can bless your sword, especially if it’s been cursed.
