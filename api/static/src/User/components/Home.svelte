<script>
  import { loadUser } from "../../helpers/user";
  import Button from "../../UI/Button.svelte";
  import EditMeal from "../../Meal/components/EditMeal.svelte";
  import config from "../../config";
  import { calculateAge } from "../../helpers/misc";

  let loading = true;
  let editMode = null;
  let user = loadUser();
  let showForm;

  let newChild = {
    name: "",
    gender: "",
    date_of_birth: ""
  };

  let children = [];
  let meals = [];

  const toggleChildForm = () => (showForm = !showForm);

  const getUserDetails = () => {
    fetch(`${config.apiUrl}user/${user.id}`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${user.token}`
      }
    })
      .then(res => {
        return res.json();
      })
      .then(res => {
        if (res.data.Children != null) children = res.data.Children;
        else children = [];
        if (res.data.Meals != null) meals = res.data.Meals;
        else meals = [];
        loading = false;
      });
  };

  const deleteChild = id => {
    fetch(`${config.apiUrl}child/${id}`, {
      method: "DELETE", // 'GET', 'PUT', 'DELETE', etc.
      headers: {
        Authorization: `Bearer ${user.token}`
      }
    }).then(res => {
      getUserDetails();
    });
  };

  const deleteMeal = id => {
    fetch(`${config.apiUrl}meal/${id}`, {
      method: "DELETE", // 'GET', 'PUT', 'DELETE', etc.
      headers: {
        Authorization: `Bearer ${user.token}`
      }
    }).then(res => {
      getUserDetails();
    });
  };

  const onSubmit = () => {
    // loading = true;
    fetch(`${config.apiUrl}child`, {
      method: "POST", // 'GET', 'PUT', 'DELETE', etc.
      headers: {
        Authorization: `Bearer ${user.token}`
      },
      body: JSON.stringify(newChild) // Coordinate the body type with 'Content-Type'
    }).then(res => {
      newChild = {
        name: "",
        gender: "",
        date_of_birth: ""
      };
      getUserDetails();
    });
  };

  const savedEdit = () => {
    editMode = null;
    getUserDetails();
  };
  const cancelEdit = () => (editMode = null);
  const startEdit = () => (editMode = "edit");

  getUserDetails();
</script>

<style>
  .user_stuff {
    display: flex;
    justify-content: space-between;
    width: 50%;
  }
</style>

{#if !loading}
  {#if editMode}
    <EditMeal on:cancel={cancelEdit} on:saved={savedEdit} {user} />
  {:else}
    <h1>Hello {user.name}</h1>
    <Button on:click={startEdit}>add meal</Button>
    <Button on:click={toggleChildForm}>Show form</Button>
    <div style={`display:${showForm || 'none'}`}>
      <div>
        <label for="name">Name</label>
        <input type="text" id="name" bind:value={newChild.name} />
      </div>
      <div>
        <label for="gender">gender</label>
        <input type="text" id="gender" bind:value={newChild.gender} />
      </div>
      <div>
        <label for="date_of_birth">date of birth</label>
        <input
          type="date"
          id="date_of_birth"
          bind:value={newChild.date_of_birth} />
      </div>
      <div>
        <Button on:click={onSubmit}>Submit</Button>
      </div>
    </div>
    <div class="user_stuff">
      <div>
        {#if children.length > 0}
          <table>
            <tr>
              <th>Name</th>
              <th>Gender</th>
              <th>Age</th>
              <th />
            </tr>
          </table>
          {#each children as child}
            <tr>
              <td>{child.name}</td>
              <td>{child.gender}</td>
              <td>{calculateAge(child.date_of_birth)}</td>
              <td>
                <button on:click={e => deleteChild(child.id)}>x</button>
              </td>
            </tr>
          {/each}
        {/if}
      </div>
      <div>
        {#if meals.length > 0}
          <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Date</th>
            <th />
          </tr>
          {#each meals as meal}
            <tr>
              <td>{meal.name}</td>
              <td>{meal.meal_type}</td>
              <td>{meal.meal_date}</td>
              <td>
                <Button color="danger" on:click={e => deleteMeal(meal.id)}>
                  X
                </Button>
                <Button href={`/meal/${meal.id}`}>Meal</Button>
              </td>

            </tr>
          {/each}
        {/if}
      </div>
    </div>
  {/if}
{/if}
