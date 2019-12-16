<script>
  import Button from "../../UI/Button.svelte";
  import config from "../../config";

  let user = {
    name: "",
    password: ""
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
        console.log(data);
        return data;
      });
    //let data = await res.json();
  };

  const onSubmit = () => {
    loginUser(user).then(res => {
      user = {
        name: "",
        password: ""
      };
      console.log(res);
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
