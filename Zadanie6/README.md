Testy funkcjonalne obejmują serwis Steam. Jako że autor nie ma wpływu na architekturę owego serwisu, niektóre testy mogą nie działać w przypadku aktualizacji Steam, błędów po stronie serwera, dużych wydarzeń (jak choćby sezonowe wyprzedaże) i zmian w przegląðarkowym front-endzie. Testy były realizowane w dniach 15-27.04, i w tych dniach były w pełni możliwe do realizacji.

Testy endpointów są zależne od stopnia zapopulowania bazy danych - powinny jednak bez problemu dać się uruchomić na bazie danych z projektu 4. Testy scenariuszy negatywnych powinny być zawsze uruchamialne i zawsze przechodzić testy poprawnie.

Testy jednostkowe są zależne wyłącznie od aplikacji, więc z ich uruchomieniem nie powinno być najmniejszych problemów.
