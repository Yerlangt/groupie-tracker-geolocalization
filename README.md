# groupie-tracker-geolocalization
## Authors: Aidana_BK, Erlangt
### Objectives

Groupie Tracker Geolocalization consists on mapping the different concerts locations of a certain artist/band given by the Client.
<a href = "https://yandex.com/dev/maps/mapsapi/"> Yandex Map Api </a> was used to get coordinates of concerts' locations (Geocoder) and to display it on the map (JS API).

**There given an API, that consists in four parts:**

- Artists: containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.
- Locations: consists in their last and/or upcoming concert locations.
- Dates: consists in their last and/or upcoming concert dates.
- Relation: does the link between all the other parts, artists, dates and locations.


**The main page has:**

- Images of the artists logo and link to the artists' personal pages

**The artist page has:**

- All information about the artist
- Map with concerts' locations

### HTTP status code
Your endpoints must return appropriate HTTP status codes.

- OK (200), if everything went without errors.
- Not Found, if nothing is found, for example templates or banners.
- Bad Request, for incorrect requests.
- Internal Server Error, for unhandled errors.

## Usage

```
After cloning the repository, complete the following commands:
    - make run
```

