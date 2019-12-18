<script>
  import config from "../config";
  import EditIngredient from "./components/EditIngredient.svelte";
  import Table from "../UI/Table.svelte";
  import { loadUser } from "../helpers/user";
  export let params = {};

  let loading = true;
  let editMode = null;
  let user = loadUser();
  let meal = {};
  let ingredients = [];

  const getMealDetails = () => {
    fetch(`${config.apiUrl}meal/${params.id}`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${user.token}`
      }
    })
      .then(res => res.json())
      .then(res => {
        meal = res.data;
        loading = false;
      });
  };

  const removeIngredient = id => {
    fetch(`${config.apiUrl}ingredient/${id}`, {
      method: "DELETE", // 'GET', 'PUT', 'DELETE', etc.
      headers: {
        Authorization: `Bearer ${user.token}`
      }
    }).then(res => {
      getMealDetails();
    });
  };

  const savedEdit = () => {
    editMode = null;
    getMealDetails();
  };
  
  const cancelEdit = () => (editMode = null);
  const startEdit = () => (editMode = "edit");

  getMealDetails();
</script>

{#if !loading}
  {#if !editMode}
    <button on:click={startEdit}>add</button>
    <p>Name : {meal.name}</p>
    <p>Type : {meal.meal_type}</p>
    <p>Date : {meal.meal_date}</p>
    {#if meal.Ingredients && meal.Ingredients.length > 0}
      <Table title="Ingredients" headers={['Ingredient', 'Calories', '']} width="50">
        {#each meal.Ingredients as ingredient}
          <tr>
            <td>{ingredient.name}</td>
            <td>{ingredient.calorie}</td>
            <td>
              <button on:click={() => removeIngredient(ingredient.id)}>
                X
              </button>
            </td>
          </tr>
        {/each}
      </Table>
    {/if}
  {:else}
    <EditIngredient
      on:cancel={cancelEdit}
      on:saved={savedEdit}
      mealID={params.id} />
  {/if}
{/if}
