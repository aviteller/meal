<script>
  import { createEventDispatcher } from "svelte";
  import Modal from "../../UI/Modal.svelte";
  import Button from "../../UI/Button.svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import config from "../../config";
  const dispatch = createEventDispatcher();

  export let id = null;
  export let user;

  let meal = {
    name: "",
    meal_type: "",
    meal_date: ""
  };

  const cancel = () => {
    dispatch("cancel");
  };

  const onSubmit = () => {
    fetch(`${config.apiUrl}meal`, {
      method: "POST", // 'GET', 'PUT', 'DELETE', etc.
      headers: {
        Authorization: `Bearer ${user.token}`
      },
      body: JSON.stringify(meal) // Coordinate the body type with 'Content-Type'
    }).then(res => {
      dispatch("saved");
    });
  };
</script>

<Modal title={id ? `Edit ${meal.name}` : 'Add new meal'} on:cancel>

  <TextInput
    id="name"
    label="Name"
    value={meal.name}
    validityMessage="Please fill in correctly"
    on:input={e => (meal.name = e.target.value)} />
  <TextInput
    id="meal_type"
    label="Meal Type"
    value={meal.meal_type}
    validityMessage="Please fill in correctly"
    on:input={e => (meal.meal_type = e.target.value)} />
  <TextInput
    id="meal_date"
    label="Date"
    type={'date'}
    value={meal.meal_date}
    validityMessage="Please fill in correctly"
    on:input={e => (meal.meal_date = e.target.value)} />

  <div slot="footer">
    <Button on:click={onSubmit}>Submit</Button>
    <Button color="danger" on:click={cancel}>Cancel</Button>
  </div>

</Modal>
