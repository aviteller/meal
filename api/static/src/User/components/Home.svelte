<script>
  import { loadUser } from "../../helpers/user";
  import Button from "../../UI/Button.svelte";
  import config from "../../config";

  let loading = true;
  let user = loadUser();

  let newChild = {
    name: "",
    gender: "",
    age: ""
  };

  let children = [];

  const getUserDetails = () => {
    fetch(`${config.apiUrl}user`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${user.token}`
      }
    })
      .then(res => {
        return res.json();
      })
      .then(res => {
        children = res.data.Children;
        loading = false;
      });
  };

  const onSubmit = () => {
    loading = true;
    fetch(`${config.apiUrl}child/new`, {
      method: "POST", // 'GET', 'PUT', 'DELETE', etc.
      headers: {
        Authorization: `Bearer ${user.token}`
      },
      body: JSON.stringify(newChild) // Coordinate the body type with 'Content-Type'
    }).then(res => {
      newChild = {
        name: "",
        gender: "",
        age: ""
      };
      getUserDetails();
    });
  };

  getUserDetails();
</script>

{#if !loading}
  <h1>Hello {user.name}</h1>
  <div>

    <div>
      <label for="name">Name</label>
      <input type="text" id="name" bind:value={newChild.name} />
    </div>
    <div>
      <label for="gender">gender</label>
      <input type="text" id="gender" bind:value={newChild.gender} />
    </div>
    <div>
      <label for="age">age</label>
      <input type="number" id="age" bind:value={newChild.age} />
    </div>
    <div>
      <Button on:click={onSubmit}>Submit</Button>
    </div>
  </div>

  {#if children.length > 0}
    {#each children as child}
      <div>{child.name} {child.gender} {child.age}</div>
    {/each}
  {/if}
{/if}
