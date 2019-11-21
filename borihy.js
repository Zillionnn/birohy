function getDaysFromBirth(birthday, date) {
    var r = Math.floor((new Date(date).getTime() - new Date(birthday).getTime()) / 1000 / 60 / 60 / 24);
    return Number(r);
}
var days = getDaysFromBirth('1993-8-7', '2019-11-21');
var phsical = days % 23;
var emotion = days % 28;
var inte = days % 33;
var state = {
    HIGH: 'HIGH PERIOD',
    LOW: 'LOW PERIOD',
    CRITICAL: 'WARNING, IN CRITICAL DAY'
};
var s = '';
if (phsical < 12) {
    s += 'phsical ';
    s += state.HIGH;
}
else if (phsical === 12) {
    s += 'phsical ';
    s += state.CRITICAL;
}
else {
    s += 'phsical ';
    s += state.LOW;
}
if (emotion < 14) {
    s += ' emotion ';
    s += state.HIGH;
}
else if (emotion === 14) {
    s += ' emotion ';
    s += state.CRITICAL;
}
else {
    s += ' emotion ';
    s += state.LOW;
}
if (inte < 17) {
    s += ' inte ';
    s += state.HIGH;
}
else if (inte === 17) {
    s += ' inte ';
    s += state.CRITICAL;
}
else {
    s += ' inte ';
    s += state.LOW;
}
console.log(phsical, emotion, inte, s);
