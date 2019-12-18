<script>
  import Button from "../../UI/Button.svelte";
  import config from "../../config";
  import Cookies from "../../helpers/Cookie";
  import { waitFor } from "../../helpers/misc";
  import { createEventDispatcher } from "svelte";
  import { push } from "svelte-spa-router";

  let c = new Cookies();
  let dispatch = createEventDispatcher();

  let user = {
    name: "",
    password: "",
    token: ""
  };

  const loginUser = async user => {
    return await fetch(`${config.apiUrl}user/login`, {
      method: "POST", // 'GET', 'PUT', 'DELETE', etc.
      body: JSON.stringify(user) // Coordinate the body type with 'Content-Type'
    })
      .then(res => {
        if (!res.ok) {
          throw new Error("Failed");
        }
        return res.json();
      })
      .then(data => {
        if (!data.status) {
          throw new Error(data.message);
        }
        return data;
      });
  };

  const onSubmit = () => {
    loginUser(user).then(res => {
      user = {
        id: res.user.id,
        name: res.user.name,
        token: res.user.token,
        loggedIn: true
      };
      c.setCookie("jwt", JSON.stringify(user), 1);
      user = {
        name: "",
        password: "",
        token: ""
      };
      waitFor(500);
      window.location.reload();
    });
  };
</script>

<h1>Login</h1>

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
