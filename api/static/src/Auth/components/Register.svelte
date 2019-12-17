<script>
  import Button from "../../UI/Button.svelte";
  import config from "../../config";
  import { push } from "svelte-spa-router";

  let user = {
    name: "",
    password: ""
  };

  const createUser = async user => {
    let res = await fetch(`${config.apiUrl}user/new`, {
      mode: "cors",
      method: "POST",
      body: JSON.stringify(user)
    });
    let data = await res.json();
    console.log(data);
    return data;
  };

  const onSubmit = () => {
    createUser(user).then(res => {
      user = {
        name: "",
        password: ""
      };
      push("/login");
    });
  };
</script>

<h1>Register</h1>

<div>
  <div>
    <label for="name">Name</label>
    <input type="text" id="name" bind:value={user.name} />
  </div>
  <div>
    <label for="password">Password</label>
    <input type="text" id="password" bind:value={user.password} />
  </div>
  <div>
    <Button on:click={onSubmit}>Submit</Button>
  </div>

</div>
