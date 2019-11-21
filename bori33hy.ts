function getDaysFromBirth(birthday: string, date: string):number {
    const r = Math.floor((new Date(date).getTime() - new Date(birthday).getTime()) / 1000 / 60 / 60 / 24)
    return r
}

let days = getDaysFromBirth('1993-8-7', '2019-11-21')
let phsical = days % 23
let emotion = days % 28
const inte = days % 33
const state = {
    HIGH: 'HIGH PERIOD',
    LOW: 'LOW PERIOD',
    CRITICAL: 'WARNING, IN CRITICAL DAY'
}
let s = ''

if (phsical < 12) {
    s += 'phsical '
    s += state.HIGH
} else if (phsical === 12) {
    s += 'phsical '
    s += state.CRITICAL
} else {
    s += 'phsical '
    s += state.LOW
}
if (emotion < 14) {
    s += ' emotion '
    s += state.HIGH
} else if (emotion === 14) {
    s += ' emotion '
    s += state.CRITICAL
} else {
    s += ' emotion '
    s += state.LOW
}
if (inte < 17) {
    s += ' inte '
    s += state.HIGH
} else if (inte === 17) {
    s += ' inte '
    s += state.CRITICAL
} else {
    s += ' inte '
    s += state.LOW
}
console.log(phsical, emotion, inte, s)