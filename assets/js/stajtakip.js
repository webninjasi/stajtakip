// https://stackoverflow.com/questions/6982692/html5-input-type-date-default-value-to-today
Date.prototype.toDateInputValue = (function() {
    var local = new Date(this);
    local.setMinutes(this.getMinutes() - this.getTimezoneOffset());
    return local.toJSON().slice(0,10);
});

$(function() {
  $('.datePicker').each(function() {
    if ($(this).val() == "") {
      $(this).val(new Date().toDateInputValue());
    }
  });
});
