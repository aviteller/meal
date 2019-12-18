export const waitFor = ms => new Promise(s => setTimeout(s, ms));
export const calculateAge = dob => {
  let firstDate = new Date(dob);
  let diff = new Date(new Date() - firstDate);

  return `${diff.toISOString().slice(0, 4) -
    1970} years and ${diff.getMonth()} months`;
};
