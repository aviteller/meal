<script>
  export let headers;
  export let width = 100;
  export let title = null
</script>

<style>
  table {
    /* width: 100%; */
    text-align: left;
  }
</style>

{#if title}
  <h1>{title}</h1>
{/if}

<table style={`width:${width}%`}>
  <thead>
    <tr>
      {#each headers as header}
        <th>{header}</th>
      {/each}
    </tr>
  </thead>
  <tbody>
    <slot />
  </tbody>
</table>
