app.component("Date_picker", {
  data() {
    return {
      startDate: (new Date()).toISOString().slice(0,10),
      endDate: (new Date()).toISOString().slice(0,10),
    };
  },
  template: `<div>
    <label for="start">Start date:</label>
    <input type="date" id="start" name="start" :value="startDate" required />

    <label for="end">End date:</label>
    <input type="date" id="end" name="end" :value="endDate" required />
  </div>
  `,
});
