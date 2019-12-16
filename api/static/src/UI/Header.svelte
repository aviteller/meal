<script>

  import Button from "./Button.svelte";
</script>

<style>
  /* header {
    position: fixed;
    width: 100%;
    top: 0;
    left: 0;
    height: 4rem;
    background-color: white;
    border-bottom: 12px solid #199b14;
    display: flex;
    justify-content: space-between;
    align-content: center;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.26);
    margin-bottom: 1rem;
  }

  h1 {
    color: black;
    margin-top: 5px;
    text-transform: uppercase;
    margin-left: 2rem; 
  }

  .links {
    margin-top: 1rem;
    margin-right: 2rem;
  } */
</style>

<header>
  <!-- <h1>
    <slot />
  </h1> -->
  <div class="links">
    <Button href="/">Home</Button>
    <Button href="/login">Login</Button>
    <Button href="/register">Register</Button>

  </div>
</header>
