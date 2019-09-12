FROM gitlab.domain.com:4567/images/debian-ssh:v1.0.0

COPY --chown=www-data:www-data tracker-sms-svc /

CMD /tracker-sms-svc
