function humanDate(t) {
    let d = Date.parse(t);

    return moment(d).format('MMM DD YYYY, HH:mm:ss')
}