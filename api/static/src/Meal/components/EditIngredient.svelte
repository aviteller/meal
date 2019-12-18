<script>
  import { createEventDispatcher } from "svelte";
  import Modal from "../../UI/Modal.svelte";
  import Button from "../../UI/Button.svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import config from "../../config";
  import { loadUser } from "../../helpers/user";
  const dispatch = createEventDispatcher();

  export let id = null;
  export let mealID;
  let user = loadUser();

  let ingredient = {
    meal_id:parseInt(mealID),
    name: "",
    calorie: ""
  };

  const cancel = () => {
    dispatch("cancel");
  };

  const onSubmit = () => {
    ingredient.calorie = parseInt(ingredient.calorie)
    fetch(`${config.apiUrl}ingredient`, {
      method: "POST", // 'GET', 'PUT', 'DELETE', etc.
      headers: {
        Authorization: `Bearer ${user.token}`
      },
      body: JSON.stringify(ingredient) // Coordinate the body type with 'Content-Type'
    }).then(res => {
      dispatch("saved");
    });
  };
</script>

<Modal title={id ? `Edit ${ingredient.name}` : 'Add new ingredient'} on:cancel>

  <TextInput
    id="name"
    label="Name"
    value={ingredient.name}
    validityMessage="Please fill in correctly"
    on:input={e => (ingredient.name = e.target.value)} />
  <TextInput
    id="meal_type"
    label="Calories"
    type="number"
    value={ingredient.calorie}
    validityMessage="Please fill in correctly"
    on:input={e => (ingredient.calorie = e.target.value)} />

  <div slot="footer">
    <Button on:click={onSubmit}>Submit</Button>
    <Button color="danger" on:click={cancel}>Cancel</Button>
  </div>

</Modal>
